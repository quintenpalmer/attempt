
host = window.location.host.split(":")[0];
port = "8888"
ws_addr = "ws://" + host + ":" + port + "/ws"
conn = new WebSocket ws_addr

# Packet IDs
LOGIN_PID = 0
PLAYER_UPDATE_PID = 1
MAP_UPDATE_PID = 2
MOVE_PLAYER_PID = 3
CHAT_PID = 4
NEARBY_PLAYER_UPDATE_PID = 5

# Packet Handling
PacketHandler = ?(Any) -> Any

mapUpdate :: PacketHandler
mapUpdate = (packet) ->
    @world.updateGrid packet.Chunks[0].Grid
    @world.dirty = true

playerUpdate :: PacketHandler
playerUpdate = (packet) ->
    @world.player.id = packet.Id
    @setPosition @world.player, packet.X, packet.Y
    @world.player.name = packet.Name
    console.log ("Player position: " + packet.X + ", " + packet.Y)
    console.log ("Camera position: " + @world.camera.getX() + ", " + @world.camera.getY())

chatUpdate :: PacketHandler
chatUpdate = (packet) ->
    console.log("Chat Packet received")
    chatString = "<" + packet.Username + "> :: " + packet.Message
    chatWindow = $('#chat_window')
    chatWindow.val(chatWindow.val() + '\n' + chatString)
    chatWindow.scrollTop(chatWindow[0].scrollHeight - chatWindow.height())

nearbyPlayerUpdate :: PacketHandler
nearbyPlayerUpdate = (packet) ->
    console.log packet.Players
    for playerPacket in packet.Players
        @world.updatePlayer playerPacket

PACKET_HANDLERS :: [...(Undefined or PacketHandler)]
PACKET_HANDLERS = (undefined for i in [0..256])

onUnloadHandler = () ->
    conn.onclose = () ->
    conn.close()

registerPacketHandler :: (Num, PacketHandler) -> Any
registerPacketHandler = (id, callback) ->
    PACKET_HANDLERS[id] = callback

initializePacketHandlers = () ->
    console.log "Initializing packet handlers..."
    registerPacketHandler(PLAYER_UPDATE_PID, playerUpdate)
    registerPacketHandler(MAP_UPDATE_PID, mapUpdate)
    registerPacketHandler(CHAT_PID, chatUpdate)
    registerPacketHandler(NEARBY_PLAYER_UPDATE_PID, nearbyPlayerUpdate)

userLogin = () ->
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

@sendMove :: (Num, Num) -> Any
@sendMove = (dx, dy) ->
    data = { Dx: dx, Dy: dy }
    sendPacket(MOVE_PLAYER_PID, data)

@sendLogin :: (Str, Str) -> Any
@sendLogin = (username, token) ->
    data = { Username: username, Token: token }
    sendPacket(LOGIN_PID, data)

@sendChat :: (Str, Str) -> Any
@sendChat = (username, message) ->
    data = { Username: username, Message: message }
    sendPacket(CHAT_PID, data)

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
    conn.onopen = (evt) ->
        userLogin()
    window.onbeforeunload = (evt) ->
        onUnloadHandler()
