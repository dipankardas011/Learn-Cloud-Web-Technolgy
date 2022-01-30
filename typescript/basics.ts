let nam: string = "Dipankar Das";

console.log(nam)


/**
 * explicitly describing
 */
interface User {
  name: string;
  id: number;
  age: number;
}

/*javascript style */
// const user = {
//   name: "Dipankar",
//   id: 23232
// };

const user: User = {
  name: "Dipankar",
  id: 324324,
  age: 34,
};

console.log(user);

class UserAccount {
  name: string;
  id: number;
  age: number;

  constructor(name: string, id: number, age: number) {
    this.name = name;
    this.id = id;
    this.age = age;
  }
}

const usr: User = new UserAccount("Dipankkkkkkk", 3432423, 34);

console.log(usr);
