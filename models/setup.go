package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://localhost:27017/?moviesPlayground"
const dbName = "movies"
const colName = "moviesCollection"

var mongoClient *mongo.Client

func Connect() {
	clientOption := options.Client().ApplyURI(connectionString)

	clientCandidate, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		panic(err)
	}

	mongoClient = clientCandidate

}
