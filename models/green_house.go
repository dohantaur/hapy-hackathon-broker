package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type GreenHouse struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Serial     string        `json:"serial" bson:"serial"`
	ActualData bson.M        `json:"actual_data" bson:"actual_data"`
	ModifiedAt time.Time     `json:"modified_at" bson:"modified_at"`
}
