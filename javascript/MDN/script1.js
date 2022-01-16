
function looping(arr) {
    for (let i of arr) {
        console.log(i);
    }
}

arr = [1, 3, 4, 5, 8, 2];

looping(arr);

// the object as human traits
let object = {
    name: "Dipankar",
    age: 34,
    gender: 'M'
};

const doubleNumbers = 1.343243234324;
console.log(doubleNumbers);
console.log(doubleNumbers.toFixed(2));
console.log(doubleNumbers.toPrecision(2));


console.log(object);
console.log(typeof object);