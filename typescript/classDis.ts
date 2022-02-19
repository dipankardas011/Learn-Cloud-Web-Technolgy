class VirtualPoint {
  x: number;
  y: number;
  constructor(x: number, y: number) {
    this.x = x;
    this.y = y;
  }
  getVal(): number {
    return this.x;
  }
}

const obj = new VirtualPoint(23, 45);
console.log(obj.getVal());
