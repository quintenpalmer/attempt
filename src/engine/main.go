package engine

import (
    "cgl.tideland.biz/applog"
    "flag"
    "net/http"
)

func Main() {
    logLevel := flag.Int("log", applog.LevelDebug, "The Logging Level[0-4]")
    flag.Parse()
    setupLogger(*logLevel)

    world.Start()
    http.HandleFunc("/ws", serveWs)
    applog.Infof("Starting websocket listener on port 8888")
    http.ListenAndServe(":8888", nil)
}

func setupLogger(level int) {
    logger := applog.GoLogger{}
    applog.SetLevel(level)
    applog.SetLogger(logger)
}
