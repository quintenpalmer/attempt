package engine

import (
    "cgl.tideland.biz/applog"
    "io"
    "time"
    "vector"
)

var (
    PLAYER_UPDATE_TIMER, _ = time.ParseDuration("500ms")
)

// TODO: Make this a non-global variable
var world = MakeWorld(10, 10)

type WorldUpdateFunc func (*World)

type World struct {
    players map[string] *Player
    worldMap *WorldMap

    register chan *connection
    unregister chan *Client
    update chan WorldUpdateFunc
}

func MakeWorld(width, height uint) *World {
    world := &World {
        make(map [string] *Player),
        MakeWorldMap(width, height),
        make(chan *connection),
        make(chan *Client),
        make(chan WorldUpdateFunc),
    }
    gm := MakeGameChunk(width, height, vector.Vector2{0, 0})
    world.worldMap.SetGameChunk(vector.Vector2{0, 0}, gm)
    return world
}

func (w *World) StartPlayerLogin(name, token string) bool {
    p, found := w.GetPlayer(name)
    if found {
        p.Login(token)
        return true
    } else {
        return false
    }
}

func (w *World) GetPlayer(name string) (*Player, bool) {
    p := w.players[name]
    if p != nil {
        return p, true
    } else {
        return nil, false
    }
}

func (w *World) Start() {
    go w.HandleConnections()
    go w.UpdateLoop()
}

func (w *World) SendPlayerInitialInfo(p *Player) {
    p.write(MakeMapPacket(w.worldMap.grid))
    p.write(MakePlayerPacket(p))
}

func (w *World) HandleConnections() {
    for conn := range w.register {
        applog.Debugf("Received request to register new player")
        conn.readTemplate(func (r io.Reader) (bool, error) {
            p := MakePlayer(5, "temp-player", vector.Vector2{0, 0})
            p.SetClient(MakeClient())
            w.players[p.name] = p
            go conn.writePump(p.client)
            go conn.readPump(p.client)
            go p.client.read()
            w.SendPlayerInitialInfo(p)
            return true, nil
        })
    }
}

func (w *World) Broadcast(payload GameWriter) {
    for _, player := range w.players {
        player.write(payload)
    }
}

func (w *World) UpdateLoop() {
    for updateFunction := range w.update {
        updateFunction(w)
    }
}
