package main

import (
	"github.com/dohantaur/hapy-hackathon-broker/application"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Println("load env")
	godotenv.Load()

	conf := application.NewConfig()
	dataStore := application.NewDataStore(conf)
	rabbit := application.NewRabbit(conf)
	r := gin.Default()

	app := application.NewApp(conf, dataStore, r, rabbit)
	app.Start()
}
