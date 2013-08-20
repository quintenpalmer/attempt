@WIDTH = 700
@HEIGHT = 400
CENTER_X = @WIDTH / 2
CENTER_Y = @HEIGHT / 2
stage = new PIXI.Stage 0xEEFFFF
renderer = PIXI.autoDetectRenderer(WIDTH, HEIGHT)
@graphics = new PIXI.Graphics()
FONT = {
    font: "14pt Arial"
    fill: "white"
    stroke: "black"
    strokeThickness: 1
}

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
        nameText = new PIXI.Text player.name, FONT
        player.sprite.addChild nameText
        nameText.position.x -= Math.floor(nameText.width / 2 - player.sprite.width / 2)
        nameText.position.y -= 20
    player.sprite.position.x = player.x + offX
    player.sprite.position.y = player.y + offY

updateSprites = () ->
    camera = @world.camera
    for p in _.values(@world.players)
        updateSprite p, CENTER_X - camera.getX(), CENTER_Y - camera.getY()

# UI CONSTANTS
UI_FONT = {
    font: "12pt Arial"
}
UI_FONT_HEALTH = {
    font: "12pt Arial"
    fill: "red"
}
UI_PLAYER_PANEL_X = 0
UI_PLAYER_PANEL_Y = 300
UI_PLAYER_PANEL_WIDTH = 200
UI_PLAYER_PANEL_HEIGHT = 100
UI_PLAYER_PANEL_BG = 0xABABAB
UI_PLAYER_PANEL_BORDER_COLOR = 0x000000
UI_PLAYER_PANEL_BORDER_SIZE = 1
UI_PLAYER_PANEL_INDENT_X = 10
UI_PLAYER_PANEL_INDENT_Y = 10
UI_PLAYER_PANEL_SEP_Y = 16

@UI = {
    drawPlayerPanel: (player) ->
        graphics.beginFill UI_PLAYER_PANEL_BG
        graphics.lineStyle UI_PLAYER_PANEL_BORDER_SIZE, UI_PLAYER_PANEL_BORDER_COLOR, 1
        graphics.drawRect UI_PLAYER_PANEL_X + 1, UI_PLAYER_PANEL_Y - 1, UI_PLAYER_PANEL_WIDTH, UI_PLAYER_PANEL_HEIGHT
        graphics.endFill()

    addText: (initText, font, x, y, parent) ->
        text = new PIXI.Text initText, font
        graphics.addChild text
        text.position.x = x
        text.position.y = y
        return text

    setupTextElements: (world) ->
        @.playerName = @.addText "", UI_FONT, UI_PLAYER_PANEL_X + UI_PLAYER_PANEL_INDENT_X, UI_PLAYER_PANEL_Y + UI_PLAYER_PANEL_INDENT_Y
        @.healthText = @.addText "", UI_FONT_HEALTH, @.playerName.position.x, @.playerName.position.y + UI_PLAYER_PANEL_SEP_Y
}

setupUI = (world) ->
    @UI.setupTextElements world

drawUI = (world) ->
    @UI.playerName.setText world.player.name
    @UI.healthText.setText (world.player.curHealth + " / " + world.player.maxHealth)
    @UI.drawPlayerPanel (world.player)

drawWorld = () ->
    updateSprites()
    xOff = CENTER_X - @world.camera.getX()
    yOff = CENTER_Y - @world.camera.getY()
    drawMap @world.grid, xOff, yOff
    drawUI @world
    requestAnimFrame animate

@startRenderer = () ->
    document.getElementById('game').appendChild renderer.view
    stage.addChild graphics
    @world.player.sprite = @stagePlayer CENTER_X, CENTER_Y
    setupUI @world
    setInterval drawWorld, 20
