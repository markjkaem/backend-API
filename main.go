package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

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
