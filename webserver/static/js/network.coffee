
host = window.location.host.split(":")[0];
port = "8888"
ws_addr = "ws://" + host + ":" + port + "/ws"
conn = new WebSocket ws_addr

# Packet IDs
LOGIN_PID = 0
PLAYER_UPDATE_PID = 1
MAP_UPDATE_PID = 2

# Packet Handling
PacketHandler = ?(Any) -> Any

mapUpdate :: PacketHandler
mapUpdate = (packet) ->
    @world.updateGrid packet.Chunks[0].Grid
    @world.dirty = true

playerUpdate :: PacketHandler
playerUpdate = (packet) ->
    @world.player.x = packet.X
    @world.player.y = packet.Y
    @world.player.name = packet.Name
    @graphics.position = new PIXI.Point x, y
    console.log("player update")


PACKET_HANDLERS :: [...(Undefined or PacketHandler)]
PACKET_HANDLERS = (undefined for i in [0..256])

registerPacketHandler :: (Num, PacketHandler) -> Any
registerPacketHandler = (id, callback) ->
    PACKET_HANDLERS[id] = callback

initializePacketHandlers = () ->
    console.log "Initializing packet handlers..."
    registerPacketHandler(PLAYER_UPDATE_PID, playerUpdate)
    registerPacketHandler(MAP_UPDATE_PID, mapUpdate)

@userLogin = () ->
    console.log "Sending login info..."
    username = $('#username').text()
    token = $('#token').text()
    sendLogin(String(username), String(token))

handlePacket :: (Str) -> Any
handlePacket = (packet) ->
    pid = packet.charCodeAt 0
    data = packet.substring 1
    console.log ("received packet " + pid + ": " + data)
    handler = PACKET_HANDLERS[pid]
    handler($.parseJSON(data))

# Packet Sending

sendLogin :: (Str, Str) -> Any
sendLogin = (username, token) ->
    data = { Username: username, Token: token }
    sendPacket(LOGIN_PID, data)

sendPacket :: (Num, Any) -> Any
sendPacket = (id, data) ->
    conn.send(String.fromCharCode(id) + JSON.stringify(data))

@sendit = () ->
    console.log(sendPacket(1, { msg: "hi" }))

@startNetworking = () ->
    console.log "Starting networking..."
    initializePacketHandlers()
    conn.onclose = (evt) ->
        console.log("Connection closed.")
    conn.onmessage = (evt) ->
        handlePacket(evt.data)
