var VirtualPoint = /** @class */ (function () {
    function VirtualPoint(x, y) {
        this.x = x;
        this.y = y;
    }
    VirtualPoint.prototype.getVal = function () {
        return this.x;
    };
    return VirtualPoint;
}());
var obj = new VirtualPoint(23, 45);
console.log(obj.getVal());
