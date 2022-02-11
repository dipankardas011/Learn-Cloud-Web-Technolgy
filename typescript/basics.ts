let nam: string = "Dipankar Das";

console.log(nam)

interface User {
  name: string;
  id: number;
  age: number;
}

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

  getID(): number {
    return this.id;
  }
  
}

const usr: UserAccount = new UserAccount("Dipankar Das", 3432423, 34);

console.log(usr);
console.log(usr.getID());


type MyBool = true | false;
type WindowStates = "open" | "closed" | "minimized";
type LockStates = "locked" | "unlocked";
type PositiveOddNumbersUnderTen = 1 | 3 | 5 | 7 | 9;


function getLength(obj: string | string[]) {
  return obj.length;
}

console.log(getLength("ABDBDBD"));
console.log(getLength(["abcd", "dcds"]));

function wrapInArray(obj: string | string[]) {
  if (typeof obj === "string") {
    return [obj];         
  }
  return obj;
}

console.log(wrapInArray("ssss"));
console.log(wrapInArray(["ssss", "sdfs", "SDfs"]));
