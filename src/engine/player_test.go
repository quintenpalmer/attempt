package engine

import (
    "unittest"
    "testing"
    "vector"
)

var (
    p = MakePlayer(0, "player0", vector.Vector2{0, 0})
)

func TestOfflineLogin(t *testing.T) {
    unittest.Check(t, p.IsOffline())
    unittest.CheckFalse(t, p.Login(""))
}

func TestOnlineLogin(t *testing.T) {
    p.loginState = ONLINE
    unittest.CheckFalse(t, p.Login(""))
}

func TestHasTokenBadLogin(t *testing.T) {
    p.loginState = TOKEN_RECEIVED
    unittest.CheckFalse(t, p.Login("bad"))
    unittest.Check(t, p.IsOffline())
}

func TestHasTokenGoodLogin(t *testing.T) {
    p.loginState = TOKEN_RECEIVED
    unittest.Check(t, p.Login(""))
    unittest.Check(t, p.IsOnline())
}