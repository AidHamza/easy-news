'use strict';
var request = require('request');
var retreiveSource = require('promise');

const API_KEY = "68406d076f664b61937e7647790cbb61";
const API_BASE = "https://newsapi.org/v1/";

module.exports = {
    loadSources: loadSources
};

function loadSources() {
  return new retreiveSource(function(fulfill, reject) {
      request.get({
          url: API_BASE + "sources",
          headers: {
            'Accept': 'application/json',
          }
      }, function(error, response, body) {
          if (!error && response.statusCode == 200) {
              var sourcesData = JSON.parse(body);
              fulfill(sourcesData);
          } else {
              console.log('-- error status :', response.statusCode, ' -- message :', error);
              reject(error);
          }
      })
  });
}