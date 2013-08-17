package engine

import (
    "cgl.tideland.biz/applog"
    "time"
    "vector"
)

var (
    PLAYER_UPDATE_TIMER, _ = time.ParseDuration("50ms")
)

// TODO: Make this a non-global variable
var world = MakeWorld(10, 10)

type WorldUpdateFunc func (*World)

type World struct {
    players map[string] *Player
    worldMap *WorldMap
    currentId uint

    register chan *connection
    unregister chan *Client
    update chan WorldUpdateFunc
}

func MakeWorld(width, height uint) *World {
    world := &World {
        make(map [string] *Player),
        MakeWorldMap(width, height),
        0,
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

func (w *World) NewPlayer(name string, token string) *Player {
    path := "players/" + name
    if res, _ := fileSystem.FileExists(path); res {
        p, err := fileSystem.LoadPlayer(path)
        if err != nil {
            applog.Errorf("Could not load player at path: %s. Error: %s",
                          path, err)
            return nil
        }
        w.players[p.Name] = p
        return p
    } else {
        applog.Debugf("Creating a new player")
        p := MakePlayer(w.getNextId(), name, vector.Vector2{0, 0})
        w.players[p.Name] = p
        return p
    }
    return nil
}

func (w *World) Start() {
    go w.HandleConnections()
    go w.UpdateLoop()
    go w.HandleShutdown()
}

func (w *World) SendPlayerInitialInfo(p *Player) {
    p.write(MakeMapPacket(w.worldMap.grid))
    p.write(MakePlayerPacket(p))
    RepeatingTimer(PLAYER_UPDATE_TIMER, func () bool {
        p.write(MakeNearbyPlayerUpdatePacket(w))
        return p.IsOnline()
    })
}

func (w *World) HandleConnections() {
    for conn := range w.register {
        applog.Debugf("Received request to register new player")
        client := MakeClient()
        go conn.writePump(client)
        go conn.readPump(client)
        go client.read()
    }
}

func (w *World) HandleShutdown() {
    for client := range w.unregister {
        applog.Debugf("Client %s logging out", client)
        path := "players/" + client.player.Name
        err := fileSystem.Save(path, client.player)
        if err != nil {
            applog.Errorf("Failed to save player at path: %s. Error: %s",
                          path, err)
        }
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

func (w *World) getNextId() uint {
    id := w.currentId
    w.currentId++
    return id
}
