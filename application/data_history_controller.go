package application

import (
	"fmt"
	"github.com/dohantaur/hapy-hackathon-broker/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

type DataHistoryController struct {
	App *App
}

func NewDataHistoryController(app *App) *DataHistoryController {
	return &DataHistoryController{App: app}
}

func (con *DataHistoryController) Post(c *gin.Context) {
	var data models.DataHistory
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad payload"})
		return
	}
	fmt.Println("**********")
	fmt.Println(data)
	fmt.Println("**********")
	col := con.App.DataStore.MongoSession.DB(con.App.Config.MongoDBName).C("data_history")
	index := mgo.Index{
		Key:        []string{"serial"},
		Unique:     false,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	col.EnsureIndex(index)

	data.CreatedAt = time.Now()
	if err := col.Insert(&data); err != nil {
		log.Fatal("cannot insert data")
	}

	d := ComputeData(data.Data)

	col = con.App.DataStore.MongoSession.DB(con.App.Config.MongoDBName).C("green_house")
	colQuerier := bson.M{"serial": &data.Serial}
	change := bson.M{"$set": bson.M{"actual_data": d, "modified_at": time.Now()}}

	err := col.Update(colQuerier, change)
	if err != nil {
		fmt.Println("cannot update")
	}
	fmt.Println(colQuerier)
	fmt.Println(change)

	c.AbortWithStatus(http.StatusCreated)
}
