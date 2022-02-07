function Person(name) {
  this.name = name;
  this.introduceSelf = function() {
    console.log(`Hi! I'm ${this.name}.`);
  }
}

const salva = new Person('Salva');
console.log(salva.name);
salva.introduceSelf();

const frankie = new Person('Frankie');
console.log(frankie.name);
frankie.introduceSelf();
