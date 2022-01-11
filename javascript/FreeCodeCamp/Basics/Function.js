
function addition(a, b) {
    console.log(a+b);
}

addition(5, 4);

// Declare the myGlobal variable below this line
let myGlobal = 10;
let oopsGlobal;
function fun1() {
    // Assign 5 to oopsGlobal Here
    oopsGlobal=5;
}

// Only change code above this line

function fun2() {
    var output = "";
    if (typeof myGlobal != "undefined") {
      output += "myGlobal: " + myGlobal;
    }
    if (typeof oopsGlobal != "undefined") {
      output += " oopsGlobal: " + oopsGlobal;
    }
    console.log(output);
}
fun1();
fun2();

let str = "wr4rfr";
let num = 234;
console.log(typeof str);
console.log(typeof num);