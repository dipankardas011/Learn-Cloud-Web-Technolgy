First, a promise can be in one of three states:

`pending`: the promise has been created, and the asynchronous function it's associated with has not succeeded or failed yet. This is the state your promise is in when it's returned from a call to fetch(), and the request is still being made.
`fulfilled`: the asynchronous function has succeeded. When a promise is fulfilled, its then() handler is called.
`rejected`: the asynchronous function has failed. When a promise is rejected, its catch() handler is called.

The promise returned by Promise.all() is:

`fulfilled` when and if all the promises in the array are fulfilled. In this case the then() handler is called with an array of all the responses, in the same order that the promises were passed into all()
`rejected` when and if any of the promises in the array are rejected. In this case the catch() handler is called with the error thrown by the promise that rejected.

Sometimes you might need any one of a set of promises to be fulfilled, and don't care which one. In that case you want Promise.any(). This is like Promise.all(), except that it is fulfilled as soon as any of the array of promises is fulfilled, or rejected if all of them are rejected:

```js
async function fetchProducts() {
  try {
    // after this line, our function will wait for the `fetch()` call to be settled
    // the `fetch()` call will either return a Response or throw an error
    const response = await fetch('https://mdn.github.io/learning-area/javascript/apis/fetching-data/can-store/products.json');
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    // after this line, our function will wait for the `response.json()` call to be settled
    // the `response.json()` call will either return the JSON object or throw an error
    const json = await response.json();
    console.log(json[0].name);
  }
  catch(error) {
    console.error(`Could not get products: ${error}`);
  }
}

fetchProducts();
```

Here we are calling await fetch(), and instead of getting a Promise, our caller gets back a fully complete Response object, just as if fetch() were a synchronous function!

We can even use a try...catch block for error handling, exactly as we would if the code were synchronous.

Note though that this magic only works inside the async function. Async functions always return a promise, so you can't do something like:
```js
async function fetchProducts() {
  try {
    const response = await fetch('https://mdn.github.io/learning-area/javascript/apis/fetching-data/can-store/products.json');
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    const json = await response.json();
    return json;
  }
  catch(error) {
    console.error(`Could not get products: ${error}`);
  }
}

const json = fetchProducts();
console.log(json[0].name);   // json is a Promise object, so this will not work
```

Instead, you'd need to do something like:

```diff
<<<<@@
async function fetchProducts() {
  try {
    const response = await fetch('https://mdn.github.io/learning-area/javascript/apis/fetching-data/can-store/products.json');
    if (!response.ok) {
      throw new Error(`HTTP error: ${response.status}`);
    }
    const json = await response.json();
    return json;
  }
  catch(error) {
   console.error(`Could not get products: ${error}`);
  }
}

const jsonPromise = fetchProducts();
- console.log(json[0].name);   // json is a Promise object, so this will not work
+ jsonPromise.then((json) => console.log(json[0].name));
```

