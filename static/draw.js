var body = document.documentElement;
var html = document.documentElement;
var shape = document.getElementById('shape');
var x = 0;
var y = 0;
var listColor = ['/green\n', '/purple\n', '/blue\n', '/pink\n', '/black\n', '/white\n', '/red\n', '/orange\n', '/brown\n', '/grey\n', '/cyan\n', '/yellow\n'];
var direction = ['/right\n', '/left\n', '/top\n', '/bot\n', '/stop\n'];
//var color = "";
var regexWidth = /width/g;
var regexHeight = /height/g;
var squareWidth = 100;
var squareHeight = 100;
var intervalMove;
var intervalDraw;
//----------Default draw----------//
shape.style.width = 100;
shape.style.height = 100;
shape.style.backgroundColor = 'black';
shape.style.position = 'absolute'; //alow to move
shape.style.top = 0;
shape.style.left = 0;
//-------------SOCKET-------------//
var socket = new WebSocket("ws://localhost:8080/msg");
console.log(socket);
//Server indicate that he's ready to receive or send data
socket.onopen = function (event) {
	//output.innerHTML += "Status: Connected\n";
};
//Receive message event from server
socket.onmessage = function (event) {
	if (event.data == '\u001b[A\n') {
		console.log(shape.style.top, y);
		y += -2;
		shape.style.top = y;
	} else if (event.data == '\u001b[B\n') {
		y += 2
		shape.style.top = y;
	} else if (event.data == '\u001b[D\n') {
		x += -2;
		shape.style.left = x;
	} else if (event.data == '\u001b[C\n') {
		x += 2;
		shape.style.left = x;
	} else {
		console.log("COMMANDE !");
		//------Transform data------//
		var elem = event.data.replace(/\//g, '');
		elem = elem.replace(/\n/g, '');
		var elements = elem.split(' ');
		var style = elements[0];
		elements.shift();
		elements = elements.join(' ');
		//------Transform end-------//
		shape.style[style] = elements; //Style
		//Save position if /top or anything..
		if (style == 'top' || style == 'bottom') {
			console.log("Y");
			y = parseInt(elements);
		} else if (style == 'left' || style == 'right') {
			console.log("X");
			x = parseInt(elements);
		}
	}
};
socket.onerror = function (event) {
	console.log("err : ", event);
}
//-----------SOCKET-END-----------//