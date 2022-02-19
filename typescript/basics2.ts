interface Point {
  x: number;
  y: number;
}

function logPoint(p: Point) {
  console.log(`${p.x}, ${p.y}`)
}

const pointer = {
  x: 12,
  y: 26
};
logPoint(pointer)

const rect = {
  x: 33,
  y: 3,
  width: 30,
  height: 80
};
logPoint(rect);

const color = {
  hex: "#3242f"
};

// logPoint(color)
