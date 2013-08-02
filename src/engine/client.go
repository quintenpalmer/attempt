package engine

import (
    "cgl.tideland.biz/applog"
    "reflect"
)

type Client struct {
    com Commander
    outgoing chan []byte
    incoming chan []byte
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
        return nil, InvalidPacketIdError(id)
    }
    err := packetStruct.UnmarshalGame(data)
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
    id := packetIds[reflect.TypeOf(payload)]
    packet := append([]byte{id}, payload.MarshalGame()...)
    applog.Debugf("Id: %d", id)
    applog.Debugf("Writing packet to client: %s", packet)
    c.outgoing<- packet
}
