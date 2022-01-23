function prodArr(arr, size) {
    if (size < 0)
        return 1;

    return prodArr(arr, size - 1) * arr[size];
}
function sum(arr, size) {
    // Only change code below this line
    if (size < 1)
        return 0;

    return sum(arr, size - 1) + arr[size - 1];
    // Only change code above this line
}
const arr = [2, 3, 5, 6];

console.log(prodArr(arr, arr.length - 1));
console.log(sum([1], 0));
console.log(sum([2, 3, 4], 1));
console.log(sum([2, 3, 4, 5], 3));
