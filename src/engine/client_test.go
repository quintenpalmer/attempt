package engine

import (
    "fmt"
    "testing"
    "unittest"
)

func TestParsePacket(t *testing.T) {
    login := LoginPacket{"hello", "world"}
    str := `{"Username":"hello","Token": "world"}`
    b := append([]byte{0}, str...)
    parsedLogin, _ := parsePacket(b)
}