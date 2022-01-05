const myConst = "Hello";
let myStr = "world ";
myStr += myConst;
console.log(myStr);

console.log("Len: %d",myStr.length);

/**
 * let myStr = "Bob";
    myStr[0] = "J";
    cannot change the value of myStr to Job, because the contents of myStr cannot be altered. Note that this does not mean that myStr cannot be changed, just that the individual characters of a string literal cannot be changed. The only way to change myStr would be to assign it with a new string, like this:

    let myStr = "Bob";
    myStr = "Job";
 */
let str = "Bob";
console.log(str);
console.log(str.charAt(0));
console.log(str[0]);
// str[0] = 'J';// this will not work
str = "Job";
console.log(str);

const myArray = ["adfa",234];
console.log(myArray);

const myArray1 = [["ABC", 34], ["DEF", 45]];
console.log(myArray1);

const arr = [2, 3, 6, 4];
console.log(arr);
arr[0] = 245;
console.log(arr);
