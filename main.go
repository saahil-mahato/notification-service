package main

import (
	"log"
	"notification-service/handlers"
	"notification-service/queue"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Inititalize the notification queue
	notificationQueue := queue.NewNotificationQueue(1, 3*time.Second)

	// Route to handle sending notifications
	app.Post("/send/:type", handlers.NotificationHandler(notificationQueue))

	// Start the worker in a separate goroutine
	go notificationQueue.StartWorker()

	log.Fatal(app.Listen(":8080"))
}
