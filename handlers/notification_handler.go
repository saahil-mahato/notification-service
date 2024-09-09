package handlers

import (
	"log"
	"notification-service/factories"
	"notification-service/services"

	"github.com/gofiber/fiber/v2"
)

// NotificationHandler handles the HTTP request to send a notification.
func NotificationHandler(c *fiber.Ctx) error {
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

	// Send the notification
	service := services.NewNotificationService(notification)
	if err := service.Send(recipient, message); err != nil {
		log.Printf("Error sending notification: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to send notification")
	}

	return c.SendString("Notification sent successfully")
}
