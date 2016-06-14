package application

import (
	"github.com/gin-gonic/gin"
	"github.com/hackathon-hapy-broker/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"time"
)

type GreenHouseController struct {
	App *App
}

func NewGreenHouseController(app *App) *GreenHouseController {
	return &GreenHouseController{App: app}
}

func (con *GreenHouseController) Post(c *gin.Context) {
	var gh models.GreenHouse
	if err := c.BindJSON(&gh); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad payload"})
		return
	}
	col := con.App.DataStore.MongoSession.DB(con.App.Config.MongoDBName).C("green_house")
	index := mgo.Index{
		Key:        []string{"serial"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	col.EnsureIndex(index)

	gh.ModifiedAt = time.Now()
	if err := col.Insert(&gh); err != nil {
		log.Fatal("cannot insert data")
	}
	c.JSON(http.StatusOK, gh)
}

func (con *GreenHouseController) One(c *gin.Context) {
	col := con.App.DataStore.MongoSession.DB(con.App.Config.MongoDBName).C("green_house")
	var data = models.GreenHouse{}
	col.Find(bson.M{"green_house_serial": c.Param("id")}).One(&data)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (con *GreenHouseController) History(c *gin.Context) {
	col := con.App.DataStore.MongoSession.DB(con.App.Config.MongoDBName).C("data_history")
	var data = []models.DataHistory{}
	start := c.DefaultQuery("start", "")
	end := c.DefaultQuery("end", "")
	if start == "" || end == "" {
		col.Find(bson.M{"serial": c.Param("id")}).All(&data)
	} else {
		col.Find(bson.M{"serial": c.Param("id"), "created_at": bson.M{"$gte": start, "$lt": end}}).All(&data)
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func (con *GreenHouseController) Action(c *gin.Context) {
	err := con.App.Rabbit.SendAction(c.DefaultQuery("name", "") + "::" + c.DefaultQuery("value", ""))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.AbortWithStatus(http.StatusNoContent)
}