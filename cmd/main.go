package main

import (
        "net/http"
        "github.com/labstack/echo"
        "github.com/AidHamza/easy-news/pkg/worker"
        "github.com/AidHamza/easy-news/pkg/sources"
        //"gopkg.in/mgo.v2"
        //"log"
)

func main() {

        sources.UpdateSources()

        go worker.GrabArticles()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

