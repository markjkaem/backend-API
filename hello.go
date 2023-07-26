// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Fiber instance
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, https://gofiber.net, http://localhost:3001, http://localhost:3000",
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	// Routes
	app.Get("/", hello)

	// Get the PORT from heroku env
	port := os.Getenv("PORT")

	// Verify if heroku provided the port or not
	if os.Getenv("PORT") == "" {
		port = "3001"
	}

	// Start server on http://${heroku-url}:${port}
	log.Fatal(app.Listen("0.0.0.0:" + port))
}


// Handler
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}