package engine

import (
    "cgl.tideland.biz/applog"
    "fmt"
)

type Packet interface {
    GameReaderWriter
    Handle(Commander)
}

type InvalidPacketIdError byte

func (ppe InvalidPacketIdError) Error() string {
    return fmt.Sprintf("Invalid packet id: %d", ppe)
}

type LoginPacket struct {
    Username string
    Token string
}

func (packet *LoginPacket) Handle(comm Commander) {
    applog.Debugf("Login request from %s with token %s",
        packet.Username, packet.Token)
}

func (packet *LoginPacket) UnmarshalGame(data []byte) error {
    return Deserialize(data, &packet)
}

func (packet *LoginPacket) MarshalGame() []byte {
    return Serialize(packet)
}

type MapPacket struct {
    Chunks []*GameChunk
}

//TODO: Get chunks visible to player
func MakeMapPacket(chunks [][]*GameChunk) *MapPacket {
    packetChunks := make([]*GameChunk, 1)
    packetChunks[0] = chunks[0][0]
    return &MapPacket{ packetChunks }
}

func (packet *MapPacket) MarshalGame() []byte {
    return Serialize(packet)
}

type PlayerPacket struct {
    X int
    Y int
    Name string
}

func MakePlayerPacket(p *Player) *PlayerPacket {
    return &PlayerPacket{ p.Position.X, p.Position.Y, p.name }
}

func (packet *PlayerPacket) MarshalGame() []byte {
    return Serialize(packet)
}

func initializePacketStructures() map[byte] Packet {
    structs := make(map [byte] Packet)
    structs[0] = new(LoginPacket)
    return structs
}

var (
    packetStructs = initializePacketStructures()
)
