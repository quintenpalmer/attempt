console.log("Hello Console!");
host = window.location.host.split(":")[0];
console.log(host);
conn = new WebSocket("ws://"+host+":8888/ws");
conn.onclose = function(evt) {
	console.log("connection closed");
}
conn.onmessage = function(evt) {
	console.log("got a message");
}

conn.send("hi");
