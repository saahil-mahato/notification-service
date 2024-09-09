package handlers

import (
	"log"
	"notification-service/factories"
	"notification-service/queue"

	"github.com/gofiber/fiber/v2"
)

// NotificationHandler handles the HTTP request to send a notification.
func NotificationHandler(notificationQueue *queue.NotificationQueue) fiber.Handler {
	return func(c *fiber.Ctx) error {
		notificationType := c.Params("type")
		recipient := c.FormValue("recipient")
		message := c.FormValue("message")

		// Use the factory to create a notification instance
		factory := &factories.NotificationFactory{}
		notification, err := factory.CreateNotification(notificationType)
		if err != nil {
			log.Printf("Error: %v\n", err)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid notification type")
		}

		// Add the notification task to the queue
		notificationQueue.AddTask(queue.NotificationTask{
			Notification: notification,
			Recipient:    recipient,
			Message:      message,
		})

		return c.SendString("Notification sent successfully")
	}
}
