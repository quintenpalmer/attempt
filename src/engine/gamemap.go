package engine

import (
    "vector"
)

type GameMap struct {
    width uint
    height uint
    grid [][] *Entity
    worldPosition vector.Vector2
}

func MakeGameMap(width, height uint, pos vector.Vector2) *GameMap {
    grid := make([][] *Entity, width)
    for x := range grid {
        grid[x] = make([]*Entity, height)
    }
    return &GameMap {
        width,
        height,
        grid,
        pos,
    }
}

func (gm *GameMap) AddStaticEntity(e *Entity) {
    x, y := e.position.Values()
    gm.grid[x][y] = e
}

func (gm *GameMap) GetEntityAtPosition(v vector.Vector2) *Entity {
    x, y := v.Values()
    return gm.grid[x][y]
}
