package engine

import (
    "cgl.tideland.biz/applog"
    "encoding/json"
)

type GameReader interface {
    UnmarshalGame([]byte) error
}

type GameWriter interface {
    MarshalGame() []byte
}

type GameReaderWriter interface {
    GameReader
    GameWriter
}

func Serialize(data interface{}) []byte {
    out, err := json.Marshal(data)
    if err != nil {
        applog.Criticalf("Error serializing data: %s", err)
        panic(err)
    }
    return out
}

// out must be a pointer
func Deserialize(data []byte, out interface{}) error {
    return json.Unmarshal(data, out)
}