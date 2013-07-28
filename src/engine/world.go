package engine

import (
    "io"
    "vector"
)

// TODO: Make this a non-global variable
var world = World {
    make(map [string] *Player),
    make(chan *connection),
    make(chan *Client),
}

type World struct {
    players map[string] *Player

    register chan *connection
    unregister chan *Client
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

func (w *World) HandleConnections() {
    for conn := range w.register {
        conn.readTemplate(func (r io.Reader) (bool, error) {
            p := MakePlayer(5, "temp-player", vector.Vector2{0, 0})
            p.SetClient(MakeClient())
            w.players[p.name] = p
            conn.getClient <- p.client
            return true, nil
        })
    }
}