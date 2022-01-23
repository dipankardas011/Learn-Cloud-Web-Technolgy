// Remember that Math.random() can never quite return a 1 and, because we're rounding down, 
// it's impossible to actually get 20. This technique will give us a whole number between 0 and 19.

// Putting everything together, this is what our code looks like:

// Math.floor(Math.random() * 20);
// Create a function called randomRange that takes a range myMin and myMax and
// returns a random whole number that's greater than or equal to myMin, and is less than or equal to myMax, inclusive.
// function randomRange(myMin, myMax) {
//     return Math.floor(Math.random() * (myMax + 1 - myMin)) + myMin;
// }

// // str -> int
// // return parseInt(str);

// function convertToInteger(str) {
//     return parseInt(str, 2);
//     // str, radix
//     // radix is the base of the string
//     // 2 for the binary ynumber

// }

// console.log(convertToInteger("10011"));

// function checkSign(num) {
//     return num === 0 ? "zero" :
//         num > 0 ? "positive" : "negative";
// }

function util(arr, n) {
    if (n < 1)
        return arr
    
    arr.push(n)
    
    return util(arr, n-1)
}

function countdown(n) {
    let arr = []
    return util(arr, n)
}

console.log(countdown(10));
console.log(countdown(5));
console.log(countdown(-1));
