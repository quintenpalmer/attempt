@WIDTH = 700
@HEIGHT = 300
CENTER_X = @WIDTH / 2
CENTER_Y = @HEIGHT / 2
stage = new PIXI.Stage 0xEEFFFF
renderer = PIXI.autoDetectRenderer(WIDTH, HEIGHT)
@graphics = new PIXI.Graphics()

tileHeight = 30
tileWidth = 30

DrawTileFunc = ?(Num, Num) -> Any

makeTile :: (Num, Num, Num, Num) -> DrawTileFunc
makeTile = (bgColor, borderColor, w, h) ->
    h_2 = h / 2
    tileFunc = (x, y) ->
        graphics.beginFill bgColor
        graphics.lineStyle(1, borderColor, 1)
        graphics.moveTo(x, y)
        graphics.lineTo(x + w, y)
        graphics.lineTo(x + w, y + h)
        graphics.lineTo(x, y + h)
        graphics.lineTo(x, y)
        graphics.endFill()

# Tile Functions
grass = makeTile(0x80CF5A, 0x339900, tileWidth, tileHeight);
dirt = makeTile(0x96712F, 0x403014, tileWidth, tileHeight);
water = makeTile(0x85b9bb, 0x476263, tileWidth, tileHeight);
empty = () -> {};
tileMethods = [grass, dirt, water, empty];

animate = () ->
    renderer.render(stage)

# Any instead of [...Num] because 2d array contracts seem to be broken
drawMap :: ([...Any], Num) -> Any
drawMap = (terrain, xOffset, yOffset) ->
    drawTile :: (Num, Num) -> Any
    drawTile = (i, j) ->
        x = j * tileWidth
        y = i * tileHeight
        tileType = terrain[i][j]
        drawTileType = tileMethods[tileType]
        drawTileType(xOffset + x, yOffset + y)
    graphics.clear()
    for row, i in terrain
        for tile, j in row
            drawTile i, j
    requestAnimFrame animate

@stagePlayer :: (Num, Num) -> Any
@stagePlayer = (x, y) ->
    console.log ("Staging at " + x + "," + y)
    avatar = PIXI.Sprite.fromImage(STATIC_PREFIX + 'img/link.png')
    avatar.position.x = x
    avatar.position.y = y
    avatar.anchor.x = 0
    avatar.anchor.y = 0
    stage.addChild avatar
    return avatar

updateSprite = (player, offX, offY) ->
    if not player.sprite
        player.sprite = stagePlayer 0, 0
        player.sprite.addChild (new PIXI.Text player.name)
    player.sprite.position.x = player.x + offX
    player.sprite.position.y = player.y + offY
    console.log ("Player " + player.name + " at " + player.sprite.position.x + ", " + player.sprite.position.y)

updateSprites = () ->
    camera = @world.camera
    for p in _.values(@world.players)
        updateSprite p, CENTER_X - camera.getX(), CENTER_Y - camera.getY()

drawWorld = () ->
    updateSprites()
    xOff = CENTER_X - @world.camera.getX()
    yOff = CENTER_Y - @world.camera.getY()
    drawMap @world.grid, xOff, yOff

@startRenderer = () ->
    document.getElementById('game').appendChild renderer.view
    stage.addChild graphics
    @world.player.sprite = @stagePlayer CENTER_X, CENTER_Y
    setInterval drawWorld, 20
