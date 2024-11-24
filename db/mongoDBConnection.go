package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// File: mongoDBConnection.go
func SetupMongoDB() (*mongo.Client, context.Context, context.CancelFunc) {
	cxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(cxt, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}
	err = client.Ping(cxt, readpref.Primary())
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}

	return client, cxt, cancel
}

func getMongoDBCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("expenses-tracker").Collection(collectionName)
}

func CloseConnection(client *mongo.Client, cxt context.Context, cancel context.CancelFunc) {
	defer func() {
		cancel()
		if err := client.Disconnect(cxt); err != nil {
			panic(err)
		}
		fmt.Println("Close connection is called")
	}()
}