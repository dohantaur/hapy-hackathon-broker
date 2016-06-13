package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type DataHistory struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Serial    string        `json:"serial" bson:"serial"`
	Data      bson.M        `json:"data" bson:"data"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
}
