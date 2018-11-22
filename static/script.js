var canvas = document.getElementById('canvas');
var context = canvas.getContext('2d');
cleanCanvas()

context.lineWidth = 18;
var down = false;

canvas.addEventListener('mousemove', draw);


canvas.addEventListener('mousedown', function () {
    down = true;
    context.beginPath();
    context.moveTo(xPos, yPos);
    canvas.addEventListener('mousemove', draw)

});

canvas.addEventListener('mouseup',function () {
    down=false;
    output();
});

function draw(e) {
    xPos = e.clientX - canvas.offsetLeft;
    yPos = e.clientY - canvas.offsetTop;
    if (down == true) {
        context.lineTo(xPos, yPos);
        context.stroke();
    }
}

function cleanCanvas() {
    context.clearRect(0,0,canvas.width,canvas.height)
    context.beginPath();
    context.rect(0, 0, canvas.width, canvas.height);
    context.fillStyle = "white";
    context.fill();
}

function output() {
    var data =canvas.toDataURL();
    document.getElementById('op').value =data;
    // alert(data);
}