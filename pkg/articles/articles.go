package articles

import (
	"fmt"
	"time"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
	"errors"

	"gopkg.in/mgo.v2/bson"
	"github.com/AidHamza/easy-news/pkg/sources"
	"github.com/AidHamza/easy-news/pkg/db"
)

const API_KEY = "68406d076f664b61937e7647790cbb61"
const URL = "https://newsapi.org/v1/articles?"

type Article struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Author string `json:"author"`
	URL string `json:"url"`
	Image string `json:"urlToImage"`
	PublishTime time.Time `json:"publishedAt"`
	Source sources.Source `json:"source"`
}

type articleResponse struct {
	Status string `json:"status"`
	Source string `json:"source"`
	SortyBy string `json:"sortBy"`
	Articles []Article `json:"articles"`
}

func SaveArticle(article *Article) error {
	return nil
}

//db.articles.ensureIndex( { title: 1, author: 1 }, {unique:true} )
func Grab(sourceName string) {
	fmt.Printf("Grabbing Articles")
	articlesClient := http.Client{
        Timeout: time.Second * 5, // Maximum of 2 secs
    }

    req, err := http.NewRequest(http.MethodGet, URL + "source=" + sourceName, nil)
    if err != nil {
        log.Fatal(err)
    }

    req.Header.Set("X-Api-Key", API_KEY)

    res, getErr := articlesClient.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }

    articles := articleResponse{}

    if res.Status == "200 OK" {
        body, readErr := ioutil.ReadAll(res.Body)
	    if readErr != nil {
	        log.Fatal(readErr)
	    }

	    jsonErr := json.Unmarshal(body, &articles)
	    if jsonErr != nil {
	        log.Fatal(jsonErr)
	    }	
    } else {
    	log.Fatal(errors.New("Failed to Query News API"))
    }

    if articles.Status == "error" {
    	log.Fatal(errors.New("Error parsing Articles request"))
    }

   	dbHandler := db.Handler{}
    s, err := db.Connect("mongodb://127.0.0.1:27017")
    if err != nil {
    	panic(err)
    }
    dbHandler.Session = s
    sourceCollection := dbHandler.SetCollection("news", "sources")
    articlesCollection := dbHandler.SetCollection("news", "articles")

    sourceData := sources.Source{}
    err = sourceCollection.Find(bson.M{"slug": sourceName}).One(&sourceData)
    if err != nil {
    	log.Fatal(err)
    }

    for _, article := range articles.Articles {
    	article.Source = sourceData
    	articlesCollection.Insert(article)
    	fmt.Printf("Inserted article %+v", article)
    }
}