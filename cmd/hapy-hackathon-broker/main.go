package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hackathon-hapy-broker/application"
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
