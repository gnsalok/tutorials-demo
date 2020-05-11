promise = new Promise((resolve, reject) => {
  setTimeout(() => {
    resolve();
  }, 3000);
});

console.log(promise);

//call when it resolve
//we are using promises with AJAX request to wait to the reponse in time.
promise
  .then(() => console.log("Finally resolve!!!"))
  .then(() => {
    console.log("Im runing.");
  })
  .catch(() => {
    console.log("Something wrong!!");
  });

//called when promise reject
// promise.catch();

// Async with Promises
