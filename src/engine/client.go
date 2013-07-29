package engine

import (
    "cgl.tideland.biz/applog"
    "encoding/json"
)

type Client struct {
    com Commander
    outgoing chan []byte
    incoming chan []byte
}

type GameWriter interface {
    MarshalGame() []byte
}

type Commander interface {
    HandleCommand(map[string]interface{})
}

func MakeClient() *Client {
    return &Client{ nil, make(chan [] byte), make(chan [] byte) }
}

func (c *Client) read() {
    for packet := range c.incoming {
        applog.Debugf("Received packet from client: %s", packet)
        jsonValues := make(map[string]interface{})
        err := json.Unmarshal(packet, &jsonValues)
        if err != nil {
            applog.Criticalf("Error Unmarshalling message. %s", err)
            panic(err)
        }
        c.com.HandleCommand(jsonValues)
    }
}

func (c *Client) write(payload GameWriter) {
    packet := payload.MarshalGame()
    applog.Debugf("Writing packet to client: %s", packet)
    c.outgoing<- packet
}
