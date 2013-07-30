package engine

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