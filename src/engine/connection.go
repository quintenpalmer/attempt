package engine

import (
    "github.com/garyburd/go-websocket/websocket"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)



const (
    // Time allowed to write a message to the client.
    writeWait = 10 * time.Second

    // Time allowed to read the next message from the client.
    readWait = 60 * time.Second

    // Send pings to client with this period. Must be less than readWait.
    pingPeriod = (readWait * 9) / 10

    // Maximum message size allowed from client.
    maxMessageSize = 512
)

// connection is an middleman between the websocket connection and the hub.
type connection struct {
    // The websocket connection.
    ws *websocket.Conn

    // Lets the world return a client to this connection
    getClient chan *Client
}

// readPump pumps messages from the websocket connection to the hub.
func (c *connection) readPump(client *Client) {
    defer func() {
        world.unregister <- client
        c.ws.Close()
    }()
    c.ws.SetReadLimit(maxMessageSize)
    c.ws.SetReadDeadline(time.Now().Add(readWait))
    c.readTemplate(func (r io.Reader) (bool, error) {
        message, err := ioutil.ReadAll(r)
        if err == nil {
            client.incoming <- message
        }
        return false, err
    })
}

func (c *connection) readTemplate(handler func(io.Reader) (bool, error)) {
    for {
        op, r, err := c.ws.NextReader()
        if err != nil {
            break
        }
        switch op {
        case websocket.OpPong:
            c.ws.SetReadDeadline(time.Now().Add(readWait))
        case websocket.OpText:
            if quit, err := handler(r); quit || err != nil {
                break
            }
        }
    }
}

// write writes a message with the given opCode and payload.
func (c *connection) write(opCode int, payload []byte) error {
    c.ws.SetWriteDeadline(time.Now().Add(writeWait))
    return c.ws.WriteMessage(opCode, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (c *connection) writePump(client *Client) {
    ticker := time.NewTicker(pingPeriod)
    defer func() {
        ticker.Stop()
        c.ws.Close()
    }()
    for {
        select {
        case message, ok := <-client.outgoing:
            if !ok {
                c.write(websocket.OpClose, []byte{})
                return
            }
            if err := c.write(websocket.OpText, message); err != nil {
                return
            }
        case <-ticker.C:
            if err := c.write(websocket.OpPing, []byte{}); err != nil {
                return
            }
        }
    }
}

// serverWs handles webocket requests from the client.
func serveWs(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", 405)
        return
    }
    if r.Header.Get("Origin") != "http://"+r.Host {
        http.Error(w, "Origin not allowed", 403)
        return
    }
    ws, err := websocket.Upgrade(w, r.Header, nil, 1024, 1024)
    if _, ok := err.(websocket.HandshakeError); ok {
        http.Error(w, "Not a websocket handshake", 400)
        return
    } else if err != nil {
        log.Println(err)
        return
    }
    c := &connection {
        ws: ws,
        getClient: make(chan *Client),
    }
    // Register with the world. Don't start the read/write loops until we get
    // a client back since the world will read from the websocket to get
    // authentication information.
    world.register <- c
    client := <-c.getClient
    go c.writePump(client)
    c.readPump(client)
}