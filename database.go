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

func querydb(name string) []map[string]interface{} {

	client := connectdb()
	if client == nil {
		log.Fatal("MongoDB client is nil")
	}
	collection := client.Database("main").Collection(name)

	cur, err := collection.Find(context.Background(), bson.D{})

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}

	var response []map[string]interface{}

	for cur.Next(context.Background()) {

		raw := cur.Current

		resultMap := make(map[string]interface{})

		if err := bson.Unmarshal(raw, &resultMap); err != nil {
			log.Fatal(err)
		}

		response = append(response, resultMap)
	}
	return response
}
