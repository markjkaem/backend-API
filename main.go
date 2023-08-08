package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://nuxt-go-three.vercel.app, http://localhost:3000, https://drizzleorm.vercel.app, https://babachulz.vercel.app",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/api/get-data", getData)

	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "3001"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
