package engine

import (
    "cgl.tideland.biz/applog"
)

type Client struct {
    player *Player
    outgoing chan []byte
    incoming chan []byte
}

func MakeClient() *Client {
    return &Client{ nil, make(chan [] byte), make(chan [] byte) }
}