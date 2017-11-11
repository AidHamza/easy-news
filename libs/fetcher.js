'use strict';

var source = require('../news/source');

var sourcesFetcher = function() {
var minutes = 1, the_interval = minutes * 60 * 1000;
	setInterval(function() {
		source.loadSources().then(function(sourcesPayload) {
			console.log(sourcesPayload);
			
	  	}, function(err) {
	  		console.log("Error : " + err)
	  	});

	}, the_interval);
}

module.exports = {
  sourcesFetcher: sourcesFetcher
};
