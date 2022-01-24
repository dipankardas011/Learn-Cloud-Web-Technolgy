// to find the a particular thing
const bird = ['parrot', 'lion', 'owl', 'owl'];
console.log(bird.indexOf('owl'));
console.log(bird.indexOf('cat'));
bird.push('cat');
console.log(bird);

for (const b of bird)
  console.log(b);


function adding(item) { return item + 1; }
function isValidLength(item) { return item.length > 6 && item.length < 10; }

const num = [3, 46, 6];
const newN = num.map(adding);

console.log(newN);

const city = ['London', 'Liverpool', 'Totnes', 'Edinburgh'];

console.log(city.filter(isValidLength));


const myData = 'Manchester,London,Liverpool,Birmingham,Leeds,Carlisle';
const myArray = myData.split(',');
console.log(myArray);
console.log(`length: ${myArray.length} ${myArray[0]}`);

console.log(myData.split(',').join(' '));

const products = ['Underpants:6.99',
  'Socks:5.99',
  'T-shirt:14.99',
  'Trousers:31.99',
  'Shoes:23.99'
];

let total = 0;
for (let i = 0; i < products.length; i++) {
  const breaking = (products[i].split(':'));
  const price = Number(breaking[1]);
  total += price;
  let itemText = `${breaking[0]} -- $${price}`;
  console.log(itemText);
}
console.log('Total: $' + total.toFixed(2));