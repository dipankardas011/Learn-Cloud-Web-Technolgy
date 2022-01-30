const cats = ['Leopard', 'Serval', 'Jaguar', 'Tiger', 'Caracal', 'Lion']

function setToUpperCase(string) {
  return string.toUpperCase()
}

console.log(cats)

const Ucats = cats.map(setToUpperCase)

console.log(Ucats)

function filteringCase(string) {
  return string.startsWith('L')
}

const Fcats = cats.filter(filteringCase)

console.log(Fcats)

const Fcats1 = cats.filter((iter) => {
  return iter.endsWith('l')
})

console.log(Fcats1)

/**
 * calculating sq
 */

function calculate() {
  for (let i = 0; i < 10; i++) {
    const newRes = `${i} x ${i} = ${i * i}`
    console.log(newRes)
  }
  console.log('\n..finished!')
}

// incase wondering how to clear the canvas using js
/**
 * clearBtn.addEventListener('click', ()=>{result.textContent=''})
 */

calculate()