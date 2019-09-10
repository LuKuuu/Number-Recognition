let canvas = document.getElementById('canvas');
let context = canvas.getContext('2d');
cleanCanvas();

let drawing = false;
context.lineWidth = 18;

//for mouse
function mouseDown() {
    drawing = true;
    context.beginPath();
    context.moveTo(xPos, yPos);
}

function mouseMove(e) {
    let xPos = e.clientX - canvas.offsetLeft;
    let yPos = e.clientY - canvas.offsetTop;
    if (drawing) {
        context.lineTo(xPos, yPos);
        context.stroke();
    }
}

function mouseUp() {
    drawing = false;
    output();
}

//for touch screen
function touchStart() {
    drawing = true;
    context.beginPath();
    event.preventDefault();
}

function touchMove(e) {
    if (e.touches && e.touches.length === 1) {
        let touch = e.touches[0];
        let touchX = touch.pageX - canvas.offsetLeft;
        let touchY = touch.pageY - canvas.offsetTop;
        if (drawing) {
            context.lineTo(touchX, touchY);
            context.stroke();
        }
    }
    event.preventDefault();
}

function touchEnd() {
    console.log("touchEnd");
    drawing = false;
    output();
}

function cleanCanvas() {
    context.clearRect(0, 0, canvas.width, canvas.height)
    context.beginPath();
    context.rect(0, 0, canvas.width, canvas.height);
    context.fillStyle = "white";
    context.fill();
    document.getElementById('op').value = "null";
}

function output() {
    document.getElementById('op').value = canvas.toDataURL();
    // alert(data);
}

canvas.addEventListener('mousedown', mouseDown, false);
canvas.addEventListener('mousemove', mouseMove, false);
canvas.addEventListener('mouseup', mouseUp, false);
canvas.addEventListener('touchstart', touchStart, false);
canvas.addEventListener('touchmove', touchMove, false);
canvas.addEventListener('touchend', touchEnd, false);