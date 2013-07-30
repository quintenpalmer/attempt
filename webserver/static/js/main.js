console.log("Hello Console!");
host = window.location.host.split(":")[0];
console.log(host);
var conn;
$(document).ready(function () {
    conn = new WebSocket("ws://"+host+":8888/ws");
    conn.onclose = function(evt) {
        console.log("connection closed");
    }
    conn.onmessage = function(evt) {
        console.log("received : " + evt.data);
    }

});

// Byte String -> Boolean
function sendPacket(id, data) {
    return conn.send(String.fromCharCode(id) + data);
}

function sendit() {
    console.log(sendPacket(1, '{"msg": "hi"}'));
}
