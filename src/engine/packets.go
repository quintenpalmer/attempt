package engine

import (
    "cgl.tideland.biz/applog"
    "encoding/json"
)

type LoginPacket struct {
    Username string
    Token string
}

func (packet *LoginPacket) Handle(comm Commander) {
    applog.Debugf("Login request from %s with token %s",
        packet.Username, packet.Token)
}

func (packet *LoginPacket) Unmarshal(data []byte) error {
    err := json.Unmarshal(data, &packet)
    return err
}

func initializePacketStructures() map[byte] Packet {
    structs := make(map [byte] Packet)
    structs[0] = new(LoginPacket)
    return structs
}

var (
    packetStructs = initializePacketStructures()
)
