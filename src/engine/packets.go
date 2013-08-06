package engine

import (
    "cgl.tideland.biz/applog"
    "fmt"
    "reflect"
    "vector"
)

type Packet interface {
    GameReaderWriter
    Handle(*Player)
}

type InvalidPacketIdError byte

func (ppe InvalidPacketIdError) Error() string {
    return fmt.Sprintf("Invalid packet id: %d", ppe)
}

func InvalidPacketPanic(packet Packet) {
    id := packetIds[reflect.TypeOf(packet)]
    panic(fmt.Sprintf("Packet %d should not be sent to the server.", id))
}

type LoginPacket struct {
    Username string
    Token string
}

func (packet *LoginPacket) Handle(comm *Player) {
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

func (packet *MapPacket) UnmarshalGame(data []byte) error {
    return Deserialize(data, &packet)
}

func (packet *MapPacket) Handle(_ *Player) {
    InvalidPacketPanic(packet)
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

func (packet *PlayerPacket) UnmarshalGame(data []byte) error {
    return Deserialize(data, &packet)
}

func (packet *PlayerPacket) Handle(_ *Player) {
    InvalidPacketPanic(packet)
}

type MovePlayerPacket struct {
    Dx int
    Dy int
}

func (packet *MovePlayerPacket) MarshalGame() []byte {
    return Serialize(packet)
}

func (packet *MovePlayerPacket) UnmarshalGame(data []byte) error {
    return Deserialize(data, &packet)
}

func (packet *MovePlayerPacket) Handle(player *Player) {
    player.Move(vector.Vector2{ packet.Dx, packet.Dy })
}

func initializePacketStructures() (map[byte] Packet, map[reflect.Type] byte) {
    idsToStructs := make(map [byte] Packet)
    idsToStructs[0] = new(LoginPacket)
    idsToStructs[1] = new(PlayerPacket)
    idsToStructs[2] = new(MapPacket)
    idsToStructs[3] = new(MovePlayerPacket)
    structsToIds := make(map [reflect.Type] byte)
    for id, initStruct := range(idsToStructs) {
        tipe := reflect.TypeOf(initStruct)
        structsToIds[tipe] = id
    }
    return idsToStructs, structsToIds
}

var (
    packetStructs, packetIds = initializePacketStructures()
)
