package handlers

import (
	"notification-service/factories"
	"notification-service/queue"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// NotificationPayload is the payload structure of the API.
type NotificationPayload struct {
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}

// NotificationHandler handles the HTTP request to send a notification.
func NotificationHandler(notificationQueue *queue.NotificationQueue) fiber.Handler {
	return func(c *fiber.Ctx) error {
		notificationType := c.Params("type")

		var payload NotificationPayload
		if err := c.BodyParser(&payload); err != nil {
			logrus.Fatalf("Error parsing request body: %v", err)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
		}

		// Use the factory to create a notification instance
		factory := &factories.NotificationFactory{}
		notification, err := factory.CreateNotification(notificationType)
		if err != nil {
			logrus.Errorf("Error: %v", err)
			return c.Status(fiber.StatusBadRequest).SendString("Invalid notification type")
		}

		// Define the retry configuration
		maxRetries := 3
		retryDelay := 5 * time.Second

		// Add the notification task to the queue with retry logic
		notificationQueue.AddTask(queue.NotificationTask{
			Notification: notification,
			Recipient:    payload.Recipient,
			Message:      payload.Message,
			MaxRetries:   maxRetries,
			RetryDelay:   retryDelay,
		})

		return c.SendString("Notification queued. Will process it shortly.")
	}
}
