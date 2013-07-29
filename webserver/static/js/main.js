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
        console.log("got a message");
    }

});

function sendit() {
    console.log(conn.send('{"msg": "hi"}'));
}
