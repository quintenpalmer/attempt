var conn;
host = window.location.host.split(":")[0];

function startNetworking() {
    initializePacketHandlers();

    conn = new WebSocket("ws://"+host+":8888/ws");
    conn.onclose = function(evt) {
        console.log("connection closed");
    }
    conn.onmessage = function(evt) {
        handlePacket(evt.data)
   }
}

//---- Packet Ids
var PLAYER_UPDATE_PID = 1;
var MAP_UPDATE_PID = 2;

function mapUpdate(packet) {
    drawMap(packet.Chunks[0].Grid);
}

function playerUpdate(packet) {

}

var PACKET_HANDLERS = {};
function initializePacketHandlers() {
    PACKET_HANDLERS[MAP_UPDATE_PID] = mapUpdate;
    PACKET_HANDLERS[PLAYER_UPDATE_PID] = playerUpdate;
}

//--- Packet Senders

// String String -> Bool
function sendLogin(username, token) {
    return sendPacket(LOGIN_PID, { Username: username, Token: token })
}

// Byte String -> Bool
function sendPacket(id, data) {
    return conn.send(String.fromCharCode(id) + JSON.stringify(data));
}

function sendit() {
    draw_square(2,2);
    console.log(sendPacket(1, { msg: "hi" }));
}


//---- Packet Handlers
// String -> ?
function handlePacket(packet) {
    var pid = packet.charCodeAt(0);
    var data = packet.substring(1);
    console.log("received packet " + pid + ": " + data);
    handler = PACKET_HANDLERS[pid];
    return handler($.parseJSON(data));
}
