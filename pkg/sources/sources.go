package sources

import (
	"time"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
	"errors"

	"github.com/AidHamza/easy-news/pkg/db"
)

const URL = "https://newsapi.org/v1/sources"

type Source struct {
	Slug        string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
	SortBysAvailable []string `json:"sortBysAvailable"`
}

type SourcesResponse struct {
	Status  string `json:"status"`
	Items []Source `json:"sources"`
}

//db.sources.ensureIndex( { slug: 1, name: 1 }, {unique:true} )
func UpdateSources() {
	sourcesClient := http.Client{
        Timeout: time.Second * 5, // Maximum of 2 secs
    }

    req, err := http.NewRequest(http.MethodGet, URL, nil)
    if err != nil {
        log.Fatal(err)
    }

    res, getErr := sourcesClient.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }

    sources := SourcesResponse{}

    if res.Status == "200 OK" {
        body, readErr := ioutil.ReadAll(res.Body)
	    if readErr != nil {
	        log.Fatal(readErr)
	    }

	    jsonErr := json.Unmarshal(body, &sources)
	    if jsonErr != nil {
	        log.Fatal(jsonErr)
	    }	
    } else {
    	log.Fatal(errors.New("Failed to Query News API"))
    }

    if sources.Status == "error" {
    	log.Fatal(errors.New("Error parsing Articles request"))
    }


    dbHandler := db.Handler{}
    s, err := db.Connect("mongodb://127.0.0.1:27017")
    if err != nil {
    	panic(err)
    }
    dbHandler.Session = s

    for _, source := range sources.Items {
    	c := dbHandler.SetCollection("news", "sources")
   		c.Insert(source)
    }

}