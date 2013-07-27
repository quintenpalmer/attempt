package gameserver

import (
    "fmt"
    "net/http"
)

type GameServer struct {
    router *router
    port string
}

func NewGameServer(p string) *GameServer {
    r := newrouter()
    w := GameServer{router: r,
                   port: p}
    return &w
}

func (gs *GameServer) StartServer() {
    http.Handle("/", gs)

    fmt.Println("Starting server on port " + gs.port)
    http.ListenAndServe(gs.port, nil)
}

func (gs *GameServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Got request from path: " + r.URL.Path)
    res := gs.router.routeRequest(r)
    fmt.Fprintf(w, res.Body)
}

func (gs *GameServer) RegisterAllCallbacks(cbs map[string]interface{}) error {
    for k, v := range cbs {
        err := gs.RegisterCallback(k, v)
        if err != nil {
            return err
        }
    }

    return nil
}

func (gs *GameServer) RegisterCallback(url string, callback interface{}) error {
    err := gs.router.registerCallback(url, callback)

    return err
}
