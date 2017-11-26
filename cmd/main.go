package main

import (
	"github.com/AidHamza/easy-news/pkg/articles"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	//"github.com/AidHamza/easy-news/pkg/worker"
	"github.com/AidHamza/easy-news/pkg/sources"
	//"gopkg.in/mgo.v2"
	//"log"
)

func main() {

	//sources.UpdateSources()
	//go worker.GrabArticles()

	e := echo.New()

	v1 := e.Group("/v1")

	v1.GET("/sources", func(c echo.Context) error {
		getSources, err := sources.GetSources()
		if err != nil {
			c.JSON(http.StatusOK, err)
		}
		return c.JSON(http.StatusOK, getSources)
	})

	v1.POST("/articles/search", func(c echo.Context) error {
		queries := &articles.Search{}
		bsonQuery := bson.M{}
		if err := c.Bind(queries); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		for key, value := range queries.Query {
			bsonQuery[key] = value
		}
		//bsonQuery["author"] = bson.RegEx{"Natasha.*", ""}

		getArticles, err := articles.GetArticlesBy(bsonQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusOK, getArticles)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
