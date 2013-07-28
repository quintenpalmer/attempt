package engine

import (
    "net/http"
)

func Main() {
    go world.HandleConnections()
    http.HandleFunc("/ws", serveWs)
    http.ListenAndServe(":8888", nil)
}
