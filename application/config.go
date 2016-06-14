package application

import (
	"log"
	"os"
	"strings"
)

type Config struct {
	Port        string
	Host        string
	GinMode     string
	MongoUrl    string
	MongoDBName string
	Rabbiturl   string
}

func NewConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	mongoUrl := os.Getenv("MONGODB_URI")
	if port == "" {
		log.Fatal("MONGODB_URI must be set")
	}
	mongoDBName := os.Getenv("MONGO_DB")
	if mongoDBName == "" {
		log.Println("MONGO_DB env var not present, fallback to uri")
		splitted := strings.Split(mongoUrl, "/")
		mongoDBName = splitted[len(splitted)-1]

	}
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "release"
	}
	rabbitUrl := os.Getenv("CLOUDAMQP_URL")
	if rabbitUrl == "" {
		log.Fatal("CLOUDAMQP_URL must be set")
	}
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		log.Fatal("$REDIS_URL must be set")
	}
	host := os.Getenv("HOST")
	if host == "" {
		log.Println("$HOST not set, fallback to default")
	}
	return &Config{
		Port:        port,
		GinMode:     ginMode,
		MongoUrl:    mongoUrl,
		MongoDBName: mongoDBName,
		Rabbiturl:   rabbitUrl,
		Host:        host,
	}
}
