const myArray = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
    [[10, 11, 12], 13, 14],
];
  
console.log(myArray[3][0]);
console.log(myArray[3]);
console.log(myArray[3][0][2]);

let arr = [1, 2, 3];
arr.push(5);
console.log(arr);

const arr1 = ["Sfdf", "dsvsd"];
arr1.push(["dsfs", 42]);
console.log(arr1);
let sdf = arr1.pop();
console.log(arr1);
/**
 * @function pop() remove from the last
 * @function shift() remove from the front
 * @function push() insert to the last
 * @function unshift() insert to the front
 */
let list = [1, 2, 3, 4, 5];
console.log("Hello: ", list);
let fromFront = list.shift();
console.log("Hello: ", list);
console.log("FRom front: ", fromFront);

let list1 = [1, 34, 453, 234];
console.log(list1);
list1.shift();
list1.unshift(["sdfsf", 45]);
console.log(list1);
