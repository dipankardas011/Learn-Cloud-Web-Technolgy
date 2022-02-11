class Person {
  name;

  constructor(name) {
    this.name = name;
  }

  introduceSelf() {
    console.log(`Hi ${this.name}`);
  }
}

class Professor extends Person {
  teaches;
  constructor(name, teaches) {
    super(name);
    this.teaches = teaches;
  }
  introduceSelf() {
    console.log(`Hi ${this.name}`);
  }
  grade(paper) {
    const grade = Math.floor(Math.random() * (5 - 1) + 1);
    console.log(grade);
  }
}

class Student extends Person {
  #year

  constructor(name, year) {
    super(name);
    this.#year = year;
  }

  introduceSelf() {
    console.log(`Hi ${this.name} year ${this.#year}`)
  }
  canStudyArchery() {
    this.#somePrivateMethod()
    return this.#year > 1;
  }

  #somePrivateMethod() {
    console.log('you called me?')
  }
}

const walsh = new Professor('Walala', 'Science');
walsh.introduceSelf();

walsh.grade('my papre');

const Summers = new Student('ssss', 3);
Summers.introduceSelf();
console.log(Summers.canStudyArchery());

// Summers.#year;