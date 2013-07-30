console.log("Hello Console!");
host = window.location.host.split(":")[0];
console.log(host);
var conn;
$(document).ready(function () {
    initializePacketHandlers();

    conn = new WebSocket("ws://"+host+":8888/ws");
    conn.onclose = function(evt) {
        console.log("connection closed");
    }
    conn.onmessage = function(evt) {
        console.log("received : " + evt.data);
        //handlePacket(evt.data)
    }

});

//---- Packet Ids
var LOGIN_PID = 0;

var PACKET_HANDLERS = {};
function initializePacketHandlers() {
    //PACKET_HANDLERS[PID] = callback_function;
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
    console.log(sendPacket(1, { msg: "hi" }));
}


//---- Packet Handlers
// String -> ?
function handlePacket(packet) {
    var pid = packet.charCodeAt(0);
    var data = packet.substring(1);
    handler = PACKET_HANDLERS[pid];
    return handler(data);
}
