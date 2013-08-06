Player = ?{
    name: Str
    x: Num
    y: Num
}

TileGrid = ?([...Any])

World = ?{
    player: Player
    grid: TileGrid
    updateGrid: (TileGrid) -> Any
    dirty: Bool
}

# --- Player ---

makePlayer :: (Str, Num, Num) -> Player
makePlayer = (name, x, y) -> {
    name: name
    x: x
    y: y
}

# --- World ---

makeWorld :: (Player) -> World
makeWorld = (player) -> {
    player: player
    grid: []
    updateGrid: (tiles) ->
        this.grid = tiles
    dirty: false
}

# --- Global Exported Definitions ---

@world = makeWorld (makePlayer "test-player", 0, 0)
