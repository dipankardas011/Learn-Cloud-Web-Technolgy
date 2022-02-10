const data = {
  greet() {
    console.log(`hello world!!!! ${this.name}`)
  }
}

function Person(name) {
  this.name = name;
}
Person.prototype = data;
Person.prototype.constructor = Person;

// const carl = Object.create(data);
// We then use Object.create() to create a new object with data as its prototype
const carl = new Person("ddd");
carl.greet();
console.log(Object.hasOwn(carl, 'name'))
console.log(Object.hasOwn(carl, 'greet'))
