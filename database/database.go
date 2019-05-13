package database

import (
	"context"
	"log"
	"os"

	"cisco.com/specification-manager/logger"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	logPrefix = "MongoDB"
	esIndex   = "api_documents"
)

var db *mongo.Database
var dbHost = os.Getenv("ARCADE_MONGO_DB_HOST")
var dbUser = os.Getenv("ARCADE_MONGO_DB_USER")
var dbPass = os.Getenv("ARCADE_MONGO_DB_PASSWORD")
var dbName = os.Getenv("ARCADE_MONGO_DB_NAME")

func constructURI() string {
	uri := "mongodb://"
	if len(dbUser) > 0 && len(dbPass) > 0 {
		uri = uri + dbUser + ":" + dbPass + "@"
	}
	uri = uri + dbHost
	return uri
}

// Connect establishes a connection to MongoDB
func Connect() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(constructURI()))
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("Mongo", "Connected to database")
	db = client.Database(dbName)
}

// GetDb gets the active database connection
func GetDb() *mongo.Database {
	return db
}
