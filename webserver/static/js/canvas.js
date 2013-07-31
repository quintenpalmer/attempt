var size = 50;
var block_width;
var block_height;
var canvasName = "mainCanvas";
var canvas;
var context;

$(document).ready(function () {
	canvas = document.getElementById(canvasName);
	context = canvas.getContext("2d");
	block_width = canvas.width/size;
	block_height = canvas.height/size;
	draw_grid();
});

function draw_square(x,y) {
	context.fillStyle="#996611";
	context.fillRect(x*size,y*size,size,size);
}

function draw_grid() {
	context.lineWidth = 2;
	context.strokeStyle = "black";
	for (var i = 0; i < canvas.height; i += size) {
		context.moveTo(0,i);
		context.lineTo(canvas.width,i);
	}
	for (var i = 0; i < canvas.width; i += size) {
		context.moveTo(i,0);
		context.lineTo(i,canvas.height);
	}
	context.stroke();
}

function clearCanvas() {
	context.save();
	context.setTransform(1, 0, 0, 1, 0, 0);
	context.clearRect(0, 0, canvas.width, canvas.height);
	context.restore();
}
