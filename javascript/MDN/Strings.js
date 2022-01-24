const str1 = "Welcome";
const str2 = "to the";
const res = `${str1} ${str2} MATRIX`;
console.log(res);


let number = 32;
console.log(typeof number);
let N = number.toString();

console.log(typeof N);
console.log(typeof Number(N));

// multiple line strings
const multiLine = `Hey, I am Dipankar
nice to see you
currently learning
* javascript
* golang
* docker & kubernetes`;
console.log(multiLine);

/**
 * @string operations
 */

const Hello = "Welcome to the Matrix!";
console.log(Hello[0]);
if (Hello.includes('matrix')) {
  console.log('✅ found matrix');
} else {
  console.log('❌ found matrix');
}

console.log(Hello.toUpperCase());

if (Hello.startsWith('welcome'))
  console.log('✅ starts with welcome');
else
  console.log('❌ starts with welcome');

if (Hello.endsWith('!'))
  console.log('✅ ends with !');
else
  console.log('❌ ends with !');

// substring is slice

console.log(Hello.substring(1, 2));
console.log(Hello.substring(4));
console.log(Hello.substring(0, Hello.length - 2));

// replace
const improvedHello = Hello.replace('Matrix', 'MATRIX🎉')
console.log(Hello);
console.log(improvedHello);

const arr = {

  0: {
    present: true,
    name: ['Dipankar', 'Das'],
    college: ['KiiT'],
    marks: {
      DSA: 90,
      COA: 99,
      Maths: 80
    },
    roll: 20051554
  },
  1: {
    present: false,
    name: ['ABCD', 'DFT'],
    college: ['IIIIT'],
    marks: {
      DSA: 87,
      COA: 90,
      Maths: 100
    },
    roll: 20051324
  },
}

let primeNumbers = [2, 3, 5, 7, 11];
for (let i of primeNumbers) {
  console.log(i);
}

// console.log(arr);
for (const iter in arr) {
  const element = arr[iter];
  console.log(element);
  console.log('=======');
}
function Split() {
  const station = 'MAN675847583748sjt567654;Manchester Piccadilly';

  const code = station.slice(0, 3);
  const semiColon = station.indexOf(';');
  const name = station.slice(semiColon + 1);
  const result = `${code}: ${name}`;
  console.log(result);
  // const listItem = document.createElement('li');
  // listItem.textContent = result;
  // listItem.appendChild(listItem);
}

Split();