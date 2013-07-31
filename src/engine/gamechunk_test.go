package engine

import (
    "fmt"
    "os"
    "testing"
    "unittest"
    "vector"
)

var (
    initChunk = MakeGameChunk(10, 10, vector.Vector2{0, 0})
)

func TestMakeGameChunk(t *testing.T) {
    unittest.CheckEqual(t, initChunk.Width, uint(10))
    unittest.CheckEqual(t, initChunk.Height, uint(10))
    for x := 0; x < 10; x++ {
        for y := 0; y < 10; y++ {
            unittest.CheckEqual(t, initChunk.Grid[x][y], TILE_EMPTY)
        }
    }
}

func TestSetTile(t *testing.T) {
    v := vector.Vector2{5, 5}
    x, y := v.Values()
    initChunk.SetTileVec(v, TILE_GRASS)
    unittest.CheckEqual(t, initChunk.GetTileVec(v), TILE_GRASS)

    initChunk.SetTile(x, y, TILE_EMPTY)
    unittest.CheckEqual(t, initChunk.GetTile(x, y,), TILE_EMPTY)
}

func TestChunkSerialization(t *testing.T) {
    serial := initChunk.MarshalGame()
    f, err := os.Create("chunk.json")
    unittest.CheckNil(t, &err)
    defer func() {
        os.Remove("chunk.json")
    }()
    fmt.Fprintf(f, "%s", serial)
    f.Close()

    f2, err2 := os.Open("chunk.json")
    defer f2.Close()
    f2Stats, _ := f2.Stat()
    bytes := make([]byte, f2Stats.Size())
    f2.Read(bytes)
    unittest.CheckNil(t, &err2)
    newChunk := &GameChunk{}
    serr := newChunk.UnmarshalGame(bytes)
    unittest.CheckNil(t, &serr)

    unittest.CheckDeepEqual(t, *newChunk, *initChunk)
}
