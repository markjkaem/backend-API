package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

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
