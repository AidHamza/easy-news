package main

import (
        "fmt"
        "net/http"
        "github.com/labstack/echo"
        "github.com/AidHamza/easy-news/pkg/articles"
        "github.com/AidHamza/easy-news/pkg/sources"
        "gopkg.in/mgo.v2"
        "log"
)

func main() {
        source := sources.Source{}
        source.Name = "Al Jazeera"
        source.Category = "Fun"
        source.Lang = "FR"
        
        article := articles.Article{}
        article.Title = "Hello"
        article.Description = "World"
        article.Source = &source

        dSession, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		fmt.Printf("Error : %+v", err)
        }
        Session := dSession.Copy()
	Session.SetMode(mgo.Monotonic, true)
        c := Session.DB("news").C("articles")
        
        err = c.Insert(&article)
        if err != nil {
                log.Fatal(err)
        }

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
