// Immediate timeouts

// Using 0 as the value for setTimeout() schedules the execution of the specified callback function as soon as possible but only after the main code thread has been run.

// For instance, the code below (see it live) outputs an alert containing "Hello", then an alert containing "World" as soon as you click OK on the first alert.

setTimeout(function() {
  alert('World');
}, 0);

alert('Hello');
