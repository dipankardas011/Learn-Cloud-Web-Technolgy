let abcd = setTimeout(()=>{
  console.log('Hello world');
}, 2000);

function hello() {
  console.log("HHHHHHH!");
}

function ab(number) {
  console.log("Sqrt = ", Math.sqrt(number));
}

let abcd2 = setTimeout(hello, 4000);
clearTimeout(abcd2);
let abcd3 = setTimeout(ab, 6000, 2);
let abcd4 = setTimeout(ab, 6000, 3);

