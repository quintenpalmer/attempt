package engine

import (
    "cgl.tideland.biz/applog"
    "net/http"
)

func Main() {
    setupLogger(*logLevel)

    go world.HandleConnections()
    http.HandleFunc("/ws", serveWs)
    applog.Infof("Starting websocket listener on port 8888")
    http.ListenAndServe(":8888", nil)
}

func setupLogger(level int) {
    logger := applog.GoLogger{}
    applog.SetLevel(level)
    applog.SetLogger(logger)
}
