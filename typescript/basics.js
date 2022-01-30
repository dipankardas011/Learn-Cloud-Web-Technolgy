var nam = "Dipankar Das";
console.log(nam);
/*javascript style */
// const user = {
//   name: "Dipankar",
//   id: 23232
// };
var user = {
    name: "Dipankar",
    id: 324324,
    age: 34
};
console.log(user);
var UserAccount = /** @class */ (function () {
    function UserAccount(name, id, age) {
        this.name = name;
        this.id = id;
        this.age = age;
    }
    return UserAccount;
}());
var usr = new UserAccount("Dipankkkkkkk", 3432423, 34);
console.log(usr);
