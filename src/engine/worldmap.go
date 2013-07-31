package engine

import (
    "vector"
)

type WorldMap struct {
    width uint
    height uint
    grid [][] *GameMap
}

func MakeWorldMap(width, height uint) *WorldMap {
    grid := make([][] *GameMap, width)
    for x := range grid {
        grid[x] = make([]*GameMap, height)
    }
    return &WorldMap {
        width,
        height,
        grid,
    }
}

func (wm *WorldMap) SetGameMap(pos vector.Vector2, gm *GameMap) {
    x, y := pos.Values()
    wm.grid[x][y] = gm
}

func (wm *WorldMap) GetGameMap(pos vector.Vector2) *GameMap {
    x, y := pos.Values()
    return wm.grid[x][y]
}
