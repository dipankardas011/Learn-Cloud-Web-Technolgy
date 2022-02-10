const person = {
  name: ['Bob', 'Smith'],
  age: 32,
  bio: function() {
    console.log(`${this.name[0]} ${this.name[1]} is ${this.age} years old.`);
  },
  introduceSelf: function() {
    console.log(`Hi! I'm ${this.name[0]}.`);
  }
};

console.log(Object.getPrototypeOf(person))

person.name
person.name[0]
person.age
person.introduceSelf()
person.name[0] = 'Dipankar'
person.name[1] = 'Das'
person.bio()

person.introduceSelf = function() {
  console.log(`Hey! I'm ${this.name[0]}.`)
}

person.introduceSelf()

