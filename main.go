package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Code  string
	Price uint
}
type Person struct {
	Email   string
	Balance string
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
	message := "Mark"

	return c.SendString("Hello " + message + " 👋!")
}

func queryMessage(c *fiber.Ctx) error {
	ch := make(chan string)
	chTwo := make(chan int)

	go func() {
		result := performTask()
		ch <- result
	}()

	go func() {
		result := performTaskTwo()
		chTwo <- result
	}()

	amount := strconv.Itoa(<-chTwo)
	user := Person{
		Email:   <-ch,
		Balance: amount,
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

func performTaskTwo() int {
	return 3000
}

// db, err := gorm.Open(mysql.Open("ice7w3ai5sg3qkntzjj7:pscale_pw_CPeqjtaSekzjVPS9b1AIjLDyrSpQWxZ3N7pIWQtUUdO@tcp(aws.connect.psdb.cloud)/test?tls=true"), &gorm.Config{})

// if err != nil {
// 	panic("Failed to connect database")
// }

// db.AutoMigrate(&Item{})

// // Create
// db.Create(&Item{Code: "D42", Price: 100})

// // Read
// var product Item

// db.First(&product)

// result := map[string]interface{}{}

// db.Table("Item").Take(&result)
