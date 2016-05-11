package application

import (
	"github.com/gin-gonic/gin"
)

type App struct {
	Name      string
	Config    *Config
	DataStore *DataStore
	Router    *gin.Engine
}

func NewApp(conf *Config, dataStore *DataStore, router *gin.Engine) *App {
	return &App{
		Name:   "hackathon_hapy_broker",
		Config: conf,
		Router: router,
		DataStore: dataStore,
	}
}

func (app *App) Start() {
	Route(app)
	app.Router.Run(":" + app.Config.Port)
}
