package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectdb() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI("mongodb+srv://markteek:piKrIsraISZ1Fkmx@cluster0.jwvmkzi.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

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
