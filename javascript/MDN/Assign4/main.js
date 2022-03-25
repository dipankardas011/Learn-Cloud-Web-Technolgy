const aliceTumbling = [{
    transform: 'rotate(0) scale(1)'
  },
  {
    transform: 'rotate(512deg) scale(2)'
  },
  {
    transform: 'rotate(360deg) scale(0)'
  }
];

const aliceTiming = {
  duration: 2000,
  iterations: 1,
  fill: 'forwards'
}

const alice1 = document.querySelector("#alice1");
const alice2 = document.querySelector("#alice2");
const alice3 = document.querySelector("#alice3");

// all simultanwously
// Promise(
//   alice1.animate(aliceTumbling, aliceTiming),
//   alice2.animate(aliceTumbling, aliceTiming),
//   alice3.animate(aliceTumbling, aliceTiming)
// );

// alice1.animate(aliceTumbling, aliceTiming)
//   .addEventListener('finish', () => {
//     alice2.animate(aliceTumbling, aliceTiming)
//       .addEventListener('finish', () => {
//         alice3.animate(aliceTumbling, aliceTiming)
//           .addEventListener('finish', () => {
//             console.log('Done')
//     })
//   })
// })

async function dd() {
  const aa1 = await alice1.animate(aliceTumbling, aliceTiming).finished;
  if (!aa1.pending) {
    // console.log(`Animation 1 Done`)
    const aa2 = await alice2.animate(aliceTumbling, aliceTiming).finished;
    if (!aa2.pending) {
      // console.log(`Animation 2 Done`)
      const aa3 = await alice3.animate(aliceTumbling, aliceTiming).finished;
      // if (!aa3.pending) {
      //   console.log(`Animation 3 Done`)
      // }
    }
  }
}

dd();