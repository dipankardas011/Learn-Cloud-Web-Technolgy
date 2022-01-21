const myDog = {
    name: "Metu",
    tails: 6,
    legs: 2,
    // 2: 234, the numbers can also be properties as JS typecase it to
    // string automatically
    friends: ['Meet', 'hahah']
};
/**
 * @property of accessing @class
 * There are two ways to access the properties of an object:
 * dot notation (.) and bracket notation ([]), similar to an array.
 * Dot notation is what you use when you know the name of
 * the property you're trying to access ahead of time.
 */
console.log(myDog.friends);
console.log("My name -> ", myDog["name"])

// to dynamically access the propertyies
const Custom = {
    HelloWorld: "Welcome to the MATRIX",
    HelloUniverse: "Welcome to the Machine City"
};

function getProp(value) {
    return "Hello" + value;
}

const properties = getProp('World');
console.log(Custom[properties]);

// adding attributes by
myDog.bark = 'woof';
console.log(myDog);
delete myDog.bark;
console.log(myDog);



// use it as a dictonary but of fixed key value pair
function hello(val) {

    const Class = {
        "alpha": "Adams",
        "bravo": "Boston",
        "charlie": "Chicag",
        "delta": "Denver",
        "echo": "Easy",
        "foxtrot": "Frank"
    };
    return Class[val];
}

console.log(hello('echo'))

// Sometimes it is useful to check if the property of a given
// object exists or not.We can use the.hasOwnProperty(propname) method
// of objects to determine if that object has the given property name. .hasOwnProperty()
const myObj = {
    top: "hat",
    bottom: "pants"
};
console.log(myObj.hasOwnProperty("top"));
console.log(myObj.hasOwnProperty("hello"));
// Setup
var recordCollection = {
    2548: {
        albumTitle: 'Slippery When Wet',
        artist: 'Bon Jovi',
        tracks: ['Let It Rock', 'You Give Love a Bad Name']
    },
    2468: {
        albumTitle: '1999',
        artist: 'Prince',
        tracks: ['1999', 'Little Red Corvette']
    },
    1245: {
        artist: 'Robert Palmer',
        tracks: []
    },
    5439: {
        albumTitle: 'ABBA Gold'
    }
};

// Only change code below this line
function updateRecords(records, id, prop, value) {
    if (prop !== 'tracks' && value !== "") {
        records[id][prop] = value;
    } else if (prop === "tracks" && records[id].hasOwnProperty("tracks") === false) {
        records[id][prop] = [value];
    } else if (prop === "tracks" && value !== "") {
        records[id][prop].push(value);
    } else if (value === "") {
        delete records[id][prop];
    }
    return records;
}

updateRecords(recordCollection, 5439, 'artist', 'ABBA');
console.log(recordCollection);