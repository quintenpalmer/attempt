package engine

import (
    "cgl.tideland.biz/applog"
    "fmt"
)

type PacketParseError byte

func (ppe PacketParseError) Error() string {
    return fmt.Sprintf("Invalid packet id: %d", ppe)
}

type Packet interface {
    Handle(Commander)
    Unmarshal([]byte) error
}

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

func parsePacket(packet []byte) (*Packet, error) {
    id := packet[0]
    data := packet[1:]
    packetStruct := packetStructs[id]
    if packetStruct == nil {
        return nil, PacketParseError(id)
    }
    err := packetStruct.Unmarshal(data)
    return &packetStruct, err
}

func (c *Client) read() {
    for packet := range c.incoming {
        applog.Debugf("Received packet from client: %s", packet)
        packetStruct, err := parsePacket(packet)
        if err != nil {
            applog.Criticalf("Error Unmarshalling message. %s", err)
            continue
        }
        (*packetStruct).Handle(c.com)
    }
}

func (c *Client) write(payload GameWriter) {
    packet := payload.MarshalGame()
    applog.Debugf("Writing packet to client: %s", packet)
    c.outgoing<- packet
}
