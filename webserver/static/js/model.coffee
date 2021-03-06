
HasPosition = ?{
    x: Num
    y: Num
}

GetAttachment = ?() -> HasPosition

Camera = ?{
    x: Num
    y: Num
    getX: () -> Num
    getY: () -> Num
    attached: () -> Bool
    attach: (GetAttachment) -> Any
    detach: () -> Any
    attachment: GetAttachment or Null
}

Entity = ?{
    id: Num
    type: Num
    x: Num
    y: Num
}

Player = ?{
    id: Num
    name: Str
    x: Num
    y: Num
    curHealth: Num
    maxHealth: Num
    setHealth: (Num, Num) -> Any
    -| invariant: ->
        @.curHealth <= @.maxHealth and @.curHealth >= 0
}

TileGrid = ?([...Any])

PlayerUpdate = ?{
    Id: Num
    X: Num
    Y: Num
    Name: Str
}

World = ?{
    camera: Camera
    player: Player
    grid: TileGrid
    updateGrid: (TileGrid) -> Any
    dirty: Bool
    players: Any
    updatePlayer: (PlayerUpdate) -> Any
    entities: Any
}

# --- Player ---

makePlayer :: (Num, Str, Num, Num, Num, Num) -> Player
makePlayer = (id, name, x, y, curHealth, maxHealth) ->
    console.log ("New player: " + name + " (" + id + ")")
    p = {
        id: id
        name: name
        x: x
        y: y
        curHealth: curHealth
        maxHealth: maxHealth
        setHealth: (cur, max) ->
            this.maxHealth = max
            this.curHealth = cur
    }
    return p

world2screenX :: (Num) -> Num
world2screenX = (x) ->
    return x + @WIDTH / 2

world2screenY :: (Num) -> Num
world2screenY = (y) ->
    return y + @HEIGHT / 2

@setPosition :: (HasPosition, Num, Num) -> Any
@setPosition = (mover, newX, newY) ->
    mover.x = newX
    mover.y = newY

@setPlayerPosition :: (Player, Num, Num) -> Any
@setPlayerPosition = (p, x, y) ->
    @setPosition p, x, y
    if p.sprite
        @setPosition p.sprite.position, (world2screenX x), (world2screenY y)

# --- Camera ---
@makeCamera :: (Num, Num) -> Camera
@makeCamera = (x, y, attach) -> {
    x: x
    y: y
    attachment: null
    getX: () -> if this.attached() then this.attachment().x else x
    getY: () -> if this.attached() then this.attachment().y else y
    attached: () -> this.attachment() != null
    attach: (obj) -> this.attachment = obj
    detach: () -> this.attachment = null
}

# --- World ---

makeWorld :: (Player) -> World
makeWorld = (player) -> {
    player: player
    camera: makeCamera player.x, player.y, player
    grid: []
    updateGrid: (tiles) ->
        this.grid = tiles
    dirty: false
    players: {}
    updatePlayer: (p) ->
        id = p.Id
        if id == this.player.id
            return
        if not this.players[id]
            player = makePlayer id, p.Name, p.X, p.Y, p.CurHealth, p.MaxHealth
            this.players[id] = player
        else
            setPosition this.players[id], p.X, p.Y
            player.setHealth p.CurHealth, p.MaxHealth
    entities: {}
}

# --- Global Exported Definitions ---

@world = makeWorld (makePlayer 0, "", 0, 0, 1, 1)
@world.camera.attach (-> world.player)
