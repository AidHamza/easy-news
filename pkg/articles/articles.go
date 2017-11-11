package articles

import (
	"time"
	"gopkg.in/mgo.v2"
	"github.com/AidHamza/easy-news/pkg/db"
	"github.com/AidHamza/easy-news/pkg/sources"
)

type Article struct {
	Title string
	Description string
	Author string
	URL string
	Image string
	PublishTime time.Time
	Source *sources.Source
}

var dbC *mgo.Collection

func init() {
	dbSession := db.DB{}
	err := dbSession.Connect("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	dbC = dbSession.SetCollection("news", "articles")
}

func SaveArticle(article *Article) error {
	err := dbC.Insert(article)
	if err != nil {
		return err
	}
	return nil
}	