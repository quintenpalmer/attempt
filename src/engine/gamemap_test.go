package engine

import (
    "testing"
    "unittest"
    "vector"
)

var (
    initMap = MakeGameMap(100, 100, vector.Vector2{0, 0})
)

func TestMakeGameMap(t *testing.T) {
    unittest.CheckEqual(t, initMap.width, uint(100))
    unittest.CheckEqual(t, initMap.height, uint(100))
    for x := 0; x < 100; x++ {
        for y := 0; y < 100; y++ {
            unittest.CheckFalse(t, initMap.grid[x][y])
        }
    }
}

func TestAddStaticEntity(t *testing.T) {
    v := vector.Vector2{10, 10}
    e := &Entity{5, v}
    initMap.AddStaticEntity(e)
    unittest.CheckEqual(t, initMap.GetEntityAtPosition(v), e)
}
