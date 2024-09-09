package main

import (
	"log"
	"notification-service/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Route to handle sending notifications
	app.Post("/send/:type", handlers.NotificationHandler)

	log.Fatal(app.Listen(":8080"))
}
