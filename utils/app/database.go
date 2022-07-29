package app

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetConnection() *mongo.Database {
	if db != nil {
		return db
	}
	db = initMongo()
	return db
}

func initMongo() *mongo.Database {
	host := os.Getenv("DB_HOST")

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	database := os.Getenv("DB_DATABASE")
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority", username, password, host, database)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	return client.Database(database)
}
