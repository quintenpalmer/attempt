package engine

import (
    "testing"
    "unittest"
    "vector"
)

var (
    testWorldMap = MakeWorldMap(uint(5), uint(10))
)

func TestMakeWorldMap(t *testing.T) {
    unittest.CheckEqual(t, testWorldMap.width, uint(5))
    unittest.CheckEqual(t, testWorldMap.height, uint(10))
    for y := 0; y < 10; y++ {
        for x := 0; x < 5; x++ {
            unittest.CheckFalse(t, testWorldMap.grid[x][y])
        }
    }
}

func TestWorldMapSetGet(t *testing.T) {
    pos := vector.Vector2{0, 0}
    gm := MakeGameMap(10, 10, pos)
    testWorldMap.SetGameMap(pos, gm)
    unittest.CheckEqual(t, gm, testWorldMap.GetGameMap(pos))
}
