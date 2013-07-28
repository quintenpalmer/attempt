package engine

import (
    "net/http"
)

func Main() {
    go world.HandleConnections()
    http.HandleFunc("/ws", serveWs)
    http.ListenAndServe(":8080", nil)
}