package engine

import (
    "cgl.tideland.biz/applog"
    "fmt"
    "reflect"
    "vector"
)

type Packet interface {
    GameReaderWriter
    Handle(*World, *Player, *Client)
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

func (packet *LoginPacket) Handle(world *World, _ *Player, client *Client) {
    applog.Debugf("Login request from %s with token %s",
        packet.Username, packet.Token)
    player := world.NewPlayer(packet.Username, packet.Token)
    player.SetClient(client)
    world.SendPlayerInitialInfo(player)
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

func (packet *MapPacket) Handle(_ *World, _ *Player, _ *Client) {
    InvalidPacketPanic(packet)
}

type PlayerPacket struct {
    Id uint
    X int
    Y int
    Name string
    CurHealth uint
    MaxHealth uint
}

func MakePlayerPacket(p *Player) *PlayerPacket {
    return &PlayerPacket{
        p.Id,
        p.Position.X,
        p.Position.Y,
        p.Name,
        p.CurHealth,
        p.MaxHealth,
    }
}

func (packet *PlayerPacket) MarshalGame() []byte {
    return Serialize(packet)
}

func (packet *PlayerPacket) UnmarshalGame(data []byte) error {
    return Deserialize(data, &packet)
}

func (packet *PlayerPacket) Handle(_ *World, _ *Player, _ *Client) {
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

func (packet *MovePlayerPacket) Handle(_ *World, player *Player, _ *Client) {
    player.Move(vector.Vector2{ packet.Dx, packet.Dy })
    player.write(MakePlayerPacket(player))
}

type ChatPacket struct {
    Username string
    Message string
}

func (packet *ChatPacket) MarshalGame() []byte {
    return Serialize(packet)
}

func (packet *ChatPacket) UnmarshalGame(data []byte) error {
    return Deserialize(data, &packet)
}

func (packet *ChatPacket) Handle(_ *World, _ *Player, _ *Client) {
    world.Broadcast(packet)
}

type NearbyPlayerUpdatePacket struct {
    Players []PlayerPacket
}

func (packet *NearbyPlayerUpdatePacket) MarshalGame() []byte {
    return Serialize(packet)
}

func (packet *NearbyPlayerUpdatePacket) UnmarshalGame(data []byte) error {
    return Deserialize(data, &packet)
}

func (packet *NearbyPlayerUpdatePacket) Handle(_ *World, _ *Player, _ *Client) {
    InvalidPacketPanic(packet)
}

func MakeNearbyPlayerUpdatePacket(world *World) *NearbyPlayerUpdatePacket {
    playerUpdates := make([]PlayerPacket, 0, len(world.players))
    for _, player := range world.players {
        playerUpdates = append(playerUpdates, *MakePlayerPacket(player))
    }
    return &NearbyPlayerUpdatePacket{playerUpdates}
}

func initializePacketStructures() (map[byte] Packet, map[reflect.Type] byte) {
    idsToStructs := make(map [byte] Packet)
    idsToStructs[0] = new(LoginPacket)
    idsToStructs[1] = new(PlayerPacket)
    idsToStructs[2] = new(MapPacket)
    idsToStructs[3] = new(MovePlayerPacket)
    idsToStructs[4] = new(ChatPacket)
    idsToStructs[5] = new(NearbyPlayerUpdatePacket)
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
