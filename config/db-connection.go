package config

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectToDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetServerAPIOptions(serverAPI)
	_, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal("Error occurred while connecting to the database: ", err.Error())
	}
}
