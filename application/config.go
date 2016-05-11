package application

import (
	"log"
	"os"
)

type Config struct {
	Port        string
	GinMode     string
	MongoUrl    string
	MongoDBName string
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	mongoUrl := os.Getenv("MONGO_URL")
	if port == "" {
		log.Fatal("$MONGO_URL must be set")
	}
	mongoDBName := os.Getenv("MONGO_DB")
	if mongoDBName == "" {
		log.Fatal("MONGO_DB must be set")
	}
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "release"
	}
	return &Config{
		Port:        port,
		GinMode:     ginMode,
		MongoUrl:    mongoUrl,
		MongoDBName: mongoDBName,
	}
}
