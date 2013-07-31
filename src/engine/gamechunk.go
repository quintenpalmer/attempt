package engine

import (
    "vector"
)

type tile uint

const (
    TILE_EMPTY tile = iota
    TILE_GRASS
)

type GameChunk struct {
    Width uint
    Height uint
    Grid [][] tile
    WorldPosition vector.Vector2
}

func MakeGameChunk(width, height uint, pos vector.Vector2) *GameChunk {
    grid := make([][] tile, width)
    for x := range grid {
        grid[x] = make([]tile, height)
    }
    for y := uint(0); y < height; y++ {
        for x := uint(0); x < width; x++ {
            grid[x][y] = TILE_EMPTY
        }
    }
    return &GameChunk {
        width,
        height,
        grid,
        pos,
    }
}

func (gm *GameChunk) SetTile(x, y int, t tile) {
    gm.Grid[x][y] = t
}

func (gm *GameChunk) GetTile(x, y int) tile {
    return gm.Grid[x][y]
}

func (gm *GameChunk) SetTileVec(vec vector.Vector2, t tile) {
    x, y := vec.Values()
    gm.SetTile(x, y, t)
}

func (gm *GameChunk) GetTileVec(vec vector.Vector2) tile {
    x, y := vec.Values()
    return gm.GetTile(x, y)
}

func (gm *GameChunk) MarshalGame() []byte {
    return Serialize(gm)
}

func (gm *GameChunk) UnmarshalGame(data []byte) error {
    return Deserialize(data, gm)
}
