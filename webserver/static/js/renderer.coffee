WIDTH = 700
HEIGHT = 300
stage = new PIXI.Stage 0xEEFFFF
renderer = PIXI.autoDetectRenderer(WIDTH, HEIGHT)
graphics = new PIXI.Graphics()

tileHeight = 60
tileWidth = 60

DrawTileFunc = ?(Num, Num) -> Any

isoTile :: (Num, Num, Num, Num) -> DrawTileFunc
isoTile = (bgColor, borderColor, w, h) ->
    h_2 = h / 2
    tileFunc = (x, y) ->
        graphics.beginFill bgColor
        graphics.lineStyle(1, borderColor, 1)
        graphics.moveTo(x, y)
        graphics.lineTo(x + w, y + h_2)
        graphics.lineTo(x, y + h)
        graphics.lineTo(x - w, y + h_2)
        graphics.lineTo(x, y)
        graphics.endFill()

# Tile Functions
grass = isoTile(0x80CF5A, 0x339900, tileWidth, tileHeight);
dirt = isoTile(0x96712F, 0x403014, tileWidth, tileHeight);
water = isoTile(0x85b9bb, 0x476263, tileWidth, tileHeight);
empty = () -> {};
tileMethods = [grass, dirt, water, empty];

animate = () ->
    renderer.render(stage)

# Any instead of [...Num] because 2d array contracts seem to be broken
@drawMap :: ([...Any], Num) -> Any
@drawMap = (terrain, xOffset) ->
    drawTile :: (Num, Num) -> Any
    drawTile = (i, j) ->
        x = j * tileWidth
        y = i * tileHeight
        isoX = x - y
        isoY = (x + y) / 2
        tileType = terrain[i][j]
        drawTile = tileMethods[tileType]
        drawTile(xOffset, isoX, isoY)
    for i in [0..terrain.length-1]
        for j in [0..terrain[i].length-1]
            drawTile i, j
    requestAnimFrame animate

@startRenderer = () ->
    document.body.appendChild renderer.view
    stage.addChild graphics
