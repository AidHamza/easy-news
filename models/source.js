const Joi = require('joi');
const MongoModels = require('mongo-models');

class Source extends MongoModels {
    static create(id, name, description, category, language, country) {

      const source = {
          id,
          name,
          description,
          category,
          language,
          country
      };

      this.insertOne(source, callback);
    }

    speak() {
        console.log(`${this.id}: ${this.name}.`);
    }
}

Source.collection = 'sources'; // the mongodb collection name

Source.schema = Joi.object().keys({
    id: Joi.string().required(),
    name: Joi.string().required(),
    description: Joi.string(),
    category: Joi.string().required(),
    language: Joi.string().required(),
    country: Joi.string().required()
});

module.exports = Source;