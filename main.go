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
		AllowOrigins: "https://nuxt-go-three.vercel.app",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	app.Get("/", generateMessage)

	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "3001"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}


func generateMessage(c *fiber.Ctx) error {
	message := "Mark"
	
	return c.SendString("Hello" + message + " 👋!")
}