package application

import (
	//"fmt"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

const MAX_DATA = 10

type lastData struct {
	m    sync.Mutex
	data []bson.M
}

func (ld *lastData) Add(d bson.M) {

	ld.data = append(ld.data, d)
	if len(ld.data) >= MAX_DATA {
		ld.data = ld.data[1:]
	}
}

var ld *lastData

func ComputeData(data bson.M) bson.M {
	return data
	/*temp := 0.0
	hum := 0.0
	lum := 0.0
	moi := 0.0
	ld.m.Lock()

	ld.Add(data)

	for _, d := range ld.data {
		fmt.Println(d)
		if len(d) > 0 {
			if d["t"] != nil {
				temp = temp + d["t"].(float64)
			}
			if d["h"] != nil {
				hum = hum + d["h"].(float64)
			}
			if d["l"] != nil {
				lum = lum + d["l"].(float64)
			}
			if d["m"] != nil {
				moi = moi + d["m"].(float64)
			}
		}
	}
	ld.m.Unlock()

	if len(ld.data) > 0 {
		length := float64(len(ld.data))
		hum = hum / length
		temp = temp / length
		lum = lum / length
		moi = moi / length
	}
	fmt.Println("hum: %f", hum)
	fmt.Println("temp: %f", temp)
	fmt.Println("lum: %f", lum)
	fmt.Println("moi: %f", moi)

	computed := bson.M{
		"h": hum,
		"t": temp,
		"l": lum,
		"m": moi,
	}

	return computed*/
}

func init() {
	ld = new(lastData)
	ld.data = []bson.M{}
}
