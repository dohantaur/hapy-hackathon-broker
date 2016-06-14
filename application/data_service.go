package application

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

const MAX_DATA = 10

type lastData struct {
	m    sync.Mutex
	data []bson.M
}

func (ld *lastData) Add(d bson.M) {
	ld.m.Lock()
	ld.data = append(ld.data, d)
	if len(ld.data) >= MAX_DATA {
		ld.data = ld.data[:1]
	}
	ld.m.Unlock()
}

var ld *lastData

func ComputeData(data bson.M) bson.M {
	temp := 0.0
	hum := 0.0
	lum := 0.0
	ld.Add(data)
	ld.m.Lock()

	for _, d := range ld.data {
		fmt.Println(d)
		if len(d) > 0 {
			temp = temp + d["t"].(float64)
			hum = hum + d["h"].(float64)
			lum = lum + d["l"].(float64)
		}
	}
	ld.m.Unlock()

	if len(ld.data) > 0 {
		hum = hum / float64(len(ld.data))
		temp = temp / float64(len(ld.data))
		lum = lum / float64(len(ld.data))
	}
	fmt.Println("hum: %f", hum)
	fmt.Println("temp: %f", temp)
	fmt.Println("lum: %f", lum)

	computed := bson.M{
		"h": hum,
		"t": temp,
		"l": lum,
	}

	return computed
}

func init() {
	ld = new(lastData)
	ld.data = []bson.M{}
}