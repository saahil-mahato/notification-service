package factories

import (
	"fmt"
	"notification-service/services"
)

// NotificationFactory is responsible for creating notification instances.
type NotificationFactory struct{}

// CreateNotification creates a notification based on the type passed.
func (f *NotificationFactory) CreateNotification(notificationType string) (services.Notification, error) {
	switch notificationType {
	case "email":
		return &services.EmailNotification{}, nil
	case "sms":
		return &services.SMSNotification{}, nil
	case "push":
		return &services.PushNotification{}, nil
	default:
		return nil, fmt.Errorf("unsupported notification type: %s", notificationType)
	}
}
