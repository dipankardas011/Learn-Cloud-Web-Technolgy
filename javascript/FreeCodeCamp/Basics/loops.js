const myArray = [];

// Only change code below this line
let i = 5

while (i >= 0) {
    myArray.push(i)
    i--
}

console.log(myArray);

// // Setup
// const myArr = [2, 3, 4, 5, 6]

// // Only change code below this line
// let total = 0
// for (let i = 0; i < myArr.length; i++)
//     total += myArr[i]

// function multiplyAll(arr) {
//     let product = 1

//     for (let i = 0; i < arr.length; i++) {
//         for (let j = 0; j < arr[i].length; j++) {
//             product *= arr[i][j]
//         }
//     }
//     return product
// }

// multiplyAll([[1, 2], [3, 4], [5, 6, 7]])

// Setup
const contacts = [
    {
        firstName: "Akira",
        lastName: "Laine",
        number: "0543236543",
        likes: ["Pizza", "Coding", "Brownie Points"],
    },
    {
        firstName: "Harry",
        lastName: "Potter",
        number: "0994372684",
        likes: ["Hogwarts", "Magic", "Hagrid"],
    },
    {
        firstName: "Sherlock",
        lastName: "Holmes",
        number: "0487345643",
        likes: ["Intriguing Cases", "Violin"],
    },
    {
        firstName: "Kristian",
        lastName: "Vos",
        number: "unknown",
        likes: ["JavaScript", "Gaming", "Foxes"],
    },
];

function lookUpProfile(name, prop) {
    let flag = false;
    let ret;
    let retS="No such contact";
    contacts.forEach(element => {
        if (element["firstName"] === name) {
            flag = true;
            if (element.hasOwnProperty(prop) === false) {
                flag = false;
                retS = "No such property";
            }
            // console.log(`Hello${element[prop]}`);
            
            ret = element[prop];
        } else {
        }
    });
    if (!flag)
        return retS;
    return ret;
}

console.log(lookUpProfile("Kristian", "lastName"));
console.log(lookUpProfile("Sherlock", "likes"));
console.log(lookUpProfile("Harry", "likes"));
console.log(lookUpProfile("Bob", "number"));
console.log(lookUpProfile("Bob", "potato"));
console.log(lookUpProfile("Akira", "address"));
