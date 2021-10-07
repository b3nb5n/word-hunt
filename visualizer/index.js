// canvas setup
const canvas = document.querySelector('canvas');
const ctx = canvas.getContext('2d');
canvas.width = window.innerWidth;
canvas.height = window.innerHeight;

const armLen = Math.sqrt((canvas.width / 2) ** 2 + canvas.height ** 2) / 2;
const origin = {
	x: canvas.width / 2,
	y: canvas.height,
};

const toRadians = (angle) => angle * (Math.PI / 180);
const toDegrees = (angle) => angle * (180 / Math.PI);

// get the point a vector lands
const endpoint = (origin, angle, magnitude) => {
	res = {
		x: magnitude * Math.cos(toRadians(angle)),
		y: magnitude * -Math.sin(toRadians(angle)),
	};

	res.x += origin.x;
	res.y += origin.y;
	return res;
};

// Get the arm angles given the point to plot
const coordsToAngles = (origin, target) => {
	const a = target.x - origin.x;
	const b = target.y - origin.y;
	const hypotenuse = Math.sqrt(a ** 2 + b ** 2);
	const hypotenuseAngle = Math.asin(a / hypotenuse);

	const innerAngle = Math.acos(hypotenuse ** 2 / (2 * hypotenuse * armLen));
	const shoulderAngle = Math.PI / 2 - hypotenuseAngle + innerAngle;

	const outerAngle = Math.acos((armLen ** 2 * 2 - hypotenuse ** 2) / (armLen ** 2 * 2));
	const elbowAngle = Math.PI - outerAngle;

	return [toDegrees(shoulderAngle), toDegrees(elbowAngle)];
};

// draw the arms to the canvas
const plotPoint = (p) => {
	ctx.clearRect(0, 0, canvas.width, canvas.height);
	ctx.beginPath();

	const [v1Angle, v2Angle] = coordsToAngles(origin, p);
	const v1End = endpoint(origin, v1Angle, armLen);
	const v2End = endpoint(v1End, v1Angle - v2Angle, armLen);

	ctx.moveTo(origin.x, origin.y);
	ctx.lineTo(v1End.x, v1End.y);
	ctx.lineTo(v2End.x, v2End.y);
	ctx.stroke();
};

canvas.addEventListener('mousemove', (e) => plotPoint({ x: e.clientX, y: e.clientY }));
