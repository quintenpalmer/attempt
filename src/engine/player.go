package engine

import (
    "vector"
)

type LoginState uint


const (
    OFFLINE LoginState = iota
    TOKEN_RECEIVED
    ONLINE
)

type Player struct {
    MoveableEntity
    Named
    loginState LoginState
    Token string
    client *Client
}

func MakePlayer(id uint, name string, position vector.Vector2) *Player {
    return &Player {
        MoveableEntity{Entity{id, position}},
        Named{name},
        OFFLINE,
        "",
        nil,
    }
}

func (p *Player) SetClient(c *Client) {
    p.client = c
    c.player = p
}

func (p *Player) IsOnline() bool {
    return p.loginState == ONLINE
}

func (p *Player) IsOffline() bool {
    return p.loginState == OFFLINE
}

func (p *Player) HasToken() bool {
    return p.loginState == TOKEN_RECEIVED
}

func (p *Player) doLogin() {
    p.loginState = ONLINE
}

func (p *Player) doLogout() {
    p.loginState = OFFLINE
    p.Token = ""
}

func (p *Player) SetToken(token string) {
    p.loginState = TOKEN_RECEIVED
    p.Token = token
}

func (p *Player) Login(clientToken string) bool {
    if p.HasToken() && p.Token == clientToken {
        p.doLogin()
        return true
    } else {
        // We should never get a bad token, so reset if we do.
        p.doLogout()
        return false
    }
}

func (p *Player) MarshalGame() []byte {
    return Serialize(p)
}

func (p *Player) write(payload GameWriter) {
    p.client.write(payload)
}

func (p *Player) UnmarshalGame(data []byte) error {
    err := Deserialize(data, p)
    return err
}
