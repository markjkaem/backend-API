package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	id         string
	name       string
	collection string
	image      string
	created_at string
	updated_at string
	deleted_at string
	code       string
	price      string
}
type Person struct {
	Email   string `json:"email"`
	Balance string `json:"balance"`
}
type Candidate struct {
	Name       string   `json:"name"`
	Interests  []string `json:"interests"`
	Language   string   `json:"language"`
	Experience bool     `json:"experience"`
}

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://nuxt-go-three.vercel.app, http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/api/message", generateMessage)

	app.Get("/api/concurrency", queryMessage)

	port := os.Getenv("PORT")

	if os.Getenv("PORT") == "" {
		port = "3001"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func generateMessage(c *fiber.Ctx) error {
	// db, err := gorm.Open(mysql.Open("ice7w3ai5sg3qkntzjj7:pscale_pw_CPeqjtaSekzjVPS9b1AIjLDyrSpQWxZ3N7pIWQtUUdO@tcp(aws.connect.psdb.cloud)/test?tls=true"), &gorm.Config{})

	// if err != nil {
	// 	panic("Failed to connect database")
	// }

	// result := map[string]interface{}{}
	// db.Model(&Products{}).First(&result)

	candidates := []Candidate{
		{
			Name:       "ravi",
			Interests:  []string{"art", "coding", "music", "travel"},
			Language:   "golang",
			Experience: false,
		},
		{
			Name:       "mark",
			Interests:  []string{"art", "coding", "music", "travel"},
			Language:   "typescript",
			Experience: false,
		},
		{
			Name:       "xavi",
			Interests:  []string{"art", "coding", "music", "travel"},
			Language:   "typescript",
			Experience: false,
		},
		{
			Name:       "roger",
			Interests:  []string{"art", "coding", "music", "travel"},
			Language:   "typescript",
			Experience: false,
		},
		{
			Name:       "vascos",
			Interests:  []string{"art", "coding", "music", "travel"},
			Language:   "typescript",
			Experience: false,
		},
		{
			Name:       "ravi",
			Interests:  []string{"art", "coding", "music", "travel"},
			Language:   "golang",
			Experience: false,
		},
		{
			Name:       "mark",
			Interests:  []string{"art", "coding", "music", "travel"},
			Language:   "typescript",
			Experience: false,
		},
		{
			Name:       "xavi",
			Interests:  []string{"art", "coding", "music", "travel"},
			Language:   "typescript",
			Experience: false,
		},
		{
			Name:       "roger",
			Interests:  []string{"art", "coding", "music", "travel"},
			Language:   "typescript",
			Experience: false,
		},
		{
			Name:       "vascos",
			Interests:  []string{"art", "coding", "music", "travel"},
			Language:   "typescript",
			Experience: false,
		},
	}

	return c.JSON(fiber.Map{
		"candidates": candidates})
}

func queryMessage(c *fiber.Ctx) error {
	convertedEmail := make(chan string)
	convertedBalance := make(chan string)

	go func() {
		result := performTask()
		convertedEmail <- result
	}()

	go func() {
		result := performTaskTwo()
		convertedBalance <- result
	}()

	user := Person{
		Email:   <-convertedEmail,
		Balance: <-convertedBalance,
	}

	u, err := json.Marshal(user)

	if err != nil {
		panic(err)
	}
	return c.SendString(string(u))
}

func performTask() string {
	return "mark.teekens@outlook.com"
}

func performTaskTwo() string {
	return "10"
}
