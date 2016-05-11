package application

import (
	"gopkg.in/mgo.v2"
	"log"
)

type DataStore struct {
	MongoSession *mgo.Session
}

func NewDataStore(conf *Config) *DataStore {
	log.Println("Mongo conection")
	session, err := mgo.Dial(conf.MongoUrl)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		log.Fatal(err)
	}
	// defer session.Close()
	return &DataStore{MongoSession: session}
}
