package engine

import (
    "vector"
)

type GameChunk struct {
    Width uint
    Height uint
    Grid [][] *Entity
    WorldPosition vector.Vector2
}

func MakeGameChunk(width, height uint, pos vector.Vector2) *GameChunk {
    grid := make([][] *Entity, width)
    for x := range grid {
        grid[x] = make([]*Entity, height)
    }
    for y := uint(0); y < height; y++ {
        for x := uint(0); x < width; x++ {
            v := vector.Vector2{int(x), int(y)}
            grid[x][y] = &Entity{0, v}
        }
    }
    return &GameChunk {
        width,
        height,
        grid,
        pos,
    }
}

func (gm *GameChunk) AddStaticEntity(e *Entity) {
    x, y := e.Position.Values()
    gm.Grid[x][y] = e
}

func (gm *GameChunk) GetEntityAtPosition(v vector.Vector2) *Entity {
    x, y := v.Values()
    return gm.Grid[x][y]
}

func (gm *GameChunk) MarshalGame() []byte {
    return Serialize(gm)
}

func (gm *GameChunk) UnmarshalGame(data []byte) error {
    return Deserialize(data, gm)
}
