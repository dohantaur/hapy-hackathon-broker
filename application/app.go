package application

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type App struct {
	Name      string
	Config    *Config
	DataStore *DataStore
	Router    *gin.Engine
	Rabbit    *Rabbit
}

func NewApp(conf *Config, dataStore *DataStore, router *gin.Engine, rabbit *Rabbit) *App {
	return &App{
		Name:      "hackathon_hapy_broker",
		Config:    conf,
		Router:    router,
		DataStore: dataStore,
		Rabbit:    rabbit,
	}
}

func (app *App) Start() {
	Route(app)
	// app.Router.Run(app.Config.Host + ":" + app.Config.Port)
	endless.ListenAndServe(app.Config.Host+":"+app.Config.Port, app.Router)
}
