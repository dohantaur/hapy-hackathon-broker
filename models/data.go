package models

import "gopkg.in/mgo.v2/bson"

type Data struct {
	ID               bson.ObjectId `bson:"_id,omitempty"`
	GreenHouseSerial string `json: "green_house_serial"`
	Data             bson.M `json: "data"`
}