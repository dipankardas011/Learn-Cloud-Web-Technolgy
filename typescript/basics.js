var nam = "Dipankar Das";
console.log(nam);
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
    UserAccount.prototype.getID = function () {
        return this.id;
    };
    return UserAccount;
}());
var usr = new UserAccount("Dipankar Das", 3432423, 34);
console.log(usr);
console.log(usr.getID());
function getLength(obj) {
    return obj.length;
}
console.log(getLength("ABDBDBD"));
console.log(getLength(["abcd", "dcds"]));
function wrapInArray(obj) {
    if (typeof obj === "string") {
        return [obj];
    }
    return obj;
}
console.log(wrapInArray("ssss"));
console.log(wrapInArray(["ssss", "sdfs", "SDfs"]));
