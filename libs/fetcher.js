var newsFetcher = function() {
var minutes = 0.01, the_interval = minutes * 60 * 1000;
setInterval(function() {
  console.log("I am doing my 5 minutes check");
  // do your stuff here
}, the_interval);

console.log("Here");
}

module.exports = {
  newsFetcher: newsFetcher
};
