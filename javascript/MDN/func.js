function hello(name = "DD") {
    console.log(`Hello ${name}!`);
}

hello('ari')
hello()

const arr = [23, 45, 2]
const newA = arr.map((item) => {
    return item * item
})

console.log(newA)