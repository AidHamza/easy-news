package worker

import (
	"github.com/AidHamza/easy-news/pkg/articles"
	"time"
)

func GrabArticles() {
	for {
		<-time.After(5 * time.Second)
		go articles.Grab("techcrunch")
	}
}
