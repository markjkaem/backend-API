package main

import (
	"github.com/gofiber/fiber/v2"
)

func getData(c *fiber.Ctx) error {
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
