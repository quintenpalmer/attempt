package engine

import (
    "vector"
)

type GameChunk struct {
    width uint
    height uint
    grid [][] *Entity
    worldPosition vector.Vector2
}

func MakeGameChunk(width, height uint, pos vector.Vector2) *GameChunk {
    grid := make([][] *Entity, width)
    for x := range grid {
        grid[x] = make([]*Entity, height)
    }
    return &GameChunk {
        width,
        height,
        grid,
        pos,
    }
}

func (gm *GameChunk) AddStaticEntity(e *Entity) {
    x, y := e.position.Values()
    gm.grid[x][y] = e
}

func (gm *GameChunk) GetEntityAtPosition(v vector.Vector2) *Entity {
    x, y := v.Values()
    return gm.grid[x][y]
}
