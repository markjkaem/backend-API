package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectdb() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	connectionString := os.Getenv("MONGODB_CONNECTIONSTRING")
	if connectionString == "" {
		log.Fatalf("The connection string was not privided, error.")
	}
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	if err := client.Database("main").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	return client

}
