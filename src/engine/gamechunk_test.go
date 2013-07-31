package engine

import (
    "testing"
    "unittest"
    "vector"
)

var (
    initChunk = MakeGameChunk(100, 100, vector.Vector2{0, 0})
)

func TestMakeGameChunk(t *testing.T) {
    unittest.CheckEqual(t, initChunk.width, uint(100))
    unittest.CheckEqual(t, initChunk.height, uint(100))
    for x := 0; x < 100; x++ {
        for y := 0; y < 100; y++ {
            unittest.CheckFalse(t, initChunk.grid[x][y])
        }
    }
}

func TestAddStaticEntity(t *testing.T) {
    v := vector.Vector2{10, 10}
    e := &Entity{5, v}
    initChunk.AddStaticEntity(e)
    unittest.CheckEqual(t, initChunk.GetEntityAtPosition(v), e)
}
