const select = document.querySelector('select');
const list = document.querySelector('ul');
const h1 = document.querySelector('h1');

select.addEventListener('change', () => {
  const choice = select.value;
  let days = 0;
  switch (choice) {
    case 'January': days = 31;
      break;
    case 'February': days = 28;
      break;
    case 'March': days = 31;
      break;
    case 'April': days = 30;
      break;
    case 'May': days = 31;
      break;
    case 'June': days = 30;
      break;
    case 'July': days = 31;
      break;
    case 'August': days = 31;
      break;
    case 'September': days = 30;
      break;
    case 'October': days = 31;
      break;
    case 'November': days = 30;
      break;
    case 'December': days = 31;
      break;
  }

  createCalendar(days, choice);
});

function createCalendar(days, choice) {
  list.innerHTML = '';
  h1.textContent = choice;
  for (let i = 1; i <= days; i++) {
    const listItem = document.createElement('li');
    listItem.textContent = i;
    list.appendChild(listItem);
  }
}

createCalendar(31, 'January');
