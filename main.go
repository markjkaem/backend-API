package main

import (
	"log"
	"os"

	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://nuxt-go-three.vercel.app, http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/api/get-data", generateMessage)

	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "3001"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func generateMessage(c *fiber.Ctx) error {
	// Create channels to receive the results from Goroutines
	peopleChan := make(chan []map[string]interface{})
	ordersChan := make(chan []map[string]interface{})

	// Start Goroutines to fetch data concurrently
	go func() {
		people := querydb("people")
		peopleChan <- people
	}()
	go func() {
		orders := querydb("orders")
		ordersChan <- orders
	}()

	// Wait for the results from Goroutines
	people := <-peopleChan
	orders := <-ordersChan

	// Close the channels
	close(peopleChan)
	close(ordersChan)

	return c.JSON(fiber.Map{
		"people": people,
		"orders": orders,
	})
}

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
