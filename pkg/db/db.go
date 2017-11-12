package db

import (
	"gopkg.in/mgo.v2"
)

type Handler struct {
	Session *mgo.Session
}

func Connect(url string) (s *mgo.Session, err error)  {
	s, err = mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	s.SetMode(mgo.Monotonic, true)
	return s, nil	
}

func (d Handler) SetCollection(dbName string, collectionName string) (*mgo.Collection) {
	return d.Session.Copy().DB(dbName).C(collectionName)
}