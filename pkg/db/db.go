package db

import (
	"gopkg.in/mgo.v2"
)

type Handler struct {
	Session *mgo.Session
}

func DBHandler() *Handler {
	dbHandler := Handler{}
	s, err := connect("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	dbHandler.Session = s
	return &dbHandler
}

func connect(url string) (s *mgo.Session, err error) {
	s, err = mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	s.SetMode(mgo.Monotonic, true)
	return s, nil
}

func (d Handler) SetCollection(dbName string, collectionName string) *mgo.Collection {
	return d.Session.Copy().DB(dbName).C(collectionName)
}
