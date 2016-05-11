package application

import (
	"github.com/gin-gonic/gin"
	"github.com/hackathon-hapy-broker/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type DataController struct {
	App *App
}

func NewDataController(app *App) *DataController {
	return &DataController{App: app}
}

func (dataController *DataController) Post(c *gin.Context) {

}

func (con *DataController) All(c *gin.Context) {
	collection := con.App.DataStore.MongoSession.DB(con.App.Config.MongoDBName).C("data")
	var data = []models.Data{}
	collection.Find(bson.M{"green_house_serial": c.Param("id")}).All(&data)
	c.JSON(http.StatusOK, gin.H{"data": data})
}
