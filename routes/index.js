const Source = require('../models/source');
const source = require('../news/source');
const MongoModels = require('mongo-models');
var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
	source.loadSources().then(function(sourcesPayload) {
		//console.log(sourcesPayload);
		for (var i = sourcesPayload.sources.length - 1; i >= 0; i--) {
			console.log(sourcesPayload.sources[i]);
			Source.create(sourcesPayload.sources[i].id, sourcesPayload.sources[i].name, sourcesPayload.sources[i].description, sourcesPayload.sources[i].category, sourcesPayload.sources[i].language, sourcesPayload.sources[i].country, (err, sources) => {
		        if (err) {
		            res.status(500).json({ error: 'something blew up' });
		            return;
		        }

		        res.json(sources[0]);
		    });
		}
			
	}, function(err) {
	  		console.log("Error : " + err)
	});

	//res.render('index', { title: 'Express' });

});

module.exports = router;
