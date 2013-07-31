package engine

import (
    "vector"
)

type WorldMap struct {
    width uint
    height uint
    grid [][] *GameChunk
}

func MakeWorldMap(width, height uint) *WorldMap {
    grid := make([][] *GameChunk, width)
    for x := range grid {
        grid[x] = make([]*GameChunk, height)
    }
    return &WorldMap {
        width,
        height,
        grid,
    }
}

func (wm *WorldMap) SetGameChunk(pos vector.Vector2, gm *GameChunk) {
    x, y := pos.Values()
    wm.grid[x][y] = gm
}

func (wm *WorldMap) GetGameChunk(pos vector.Vector2) *GameChunk {
    x, y := pos.Values()
    return wm.grid[x][y]
}
