package worker

import (
	"time"
	"github.com/AidHamza/easy-news/pkg/articles"
)

func GrabArticles() {
  for {
    <-time.After(5 * time.Second)
    go articles.Grab("techcrunch")
  }
}