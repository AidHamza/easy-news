package db

import (
	"gopkg.in/mgo.v2"
)

type DB struct {
	Session *mgo.Session
}

func (d DB) Connect(url string) (err error)  {
	d.Session, err = mgo.Dial(url)
	if err != nil {
		return err
	}
	d.Session.SetMode(mgo.Monotonic, true)
	d.Session = d.Session.Copy()
	return nil	
}

func (d DB) SetCollection(dbName string, collectionName string) (*mgo.Collection) {
	return d.Session.DB(dbName).C(collectionName)
}