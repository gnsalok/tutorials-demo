promise = new Promise((resolve, reject) => {
  reject();
});

console.log(promise);

//call when it resolve
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
