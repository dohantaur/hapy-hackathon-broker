package api

/*
import (
	"github.com/hackathon-hapy-broker/application"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"github.com/hackathon-hapy-broker/models"
	"net/http"
)

type DataController struct {
	App *application.App
}

func NewDataController(app *application.App) *DataController {
	return &DataController{App: app}
}

func (dataController *DataController) Post(c *gin.Context) {

}

func (dataController *DataController) All(c *gin.Context) {
	collection := dataController.App.DataStore.MongoDB.C("data")
	data := []models.Data{}
	collection.Find(bson.M{"green_house_serial", c.Param("id")}).All(&data)
	c.JSON(http.StatusOK, gin.H{"data": data})
}*/
