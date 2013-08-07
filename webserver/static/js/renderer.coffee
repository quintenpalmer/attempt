@WIDTH = 700
@HEIGHT = 300
stage = new PIXI.Stage 0xEEFFFF
renderer = PIXI.autoDetectRenderer(WIDTH, HEIGHT)
@graphics = new PIXI.Graphics()

tileHeight = 60
tileWidth = 60

DrawTileFunc = ?(Num, Num) -> Any

makeTile :: (Num, Num, Num, Num) -> DrawTileFunc
makeTile = (bgColor, borderColor, w, h) ->
    h_2 = h / 2
    tileFunc = (x, y) ->
        console.log ("Drawing at " + x + ", " + y)
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

drawWorld = () ->
    if world.dirty
        xOff = @world.player.x + @WIDTH / 2
        yOff = @world.player.y + @HEIGHT / 2
        console.log ("coords: " + xOff + ", " + yOff)
        drawMap @world.grid, xOff, yOff


@startRenderer = () ->
    document.getElementById('game').appendChild renderer.view
    stage.addChild graphics
    setInterval drawWorld, 20
