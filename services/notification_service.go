package services

// Notification is the interface that all notification types must implement.
type Notification interface {
	Send(recipient string, message string) error
}

// NotificationService handles the delivery of notifications.
type NotificationService struct {
	strategy Notification
}

// NewNotificationService creates a new NotificationService with the specified strategy.
func NewNotificationService(strategy Notification) *NotificationService {
	return &NotificationService{strategy: strategy}
}

// Send sends the notification using the chosen strategy.
func (n *NotificationService) Send(recipient string, message string) error {
	return n.strategy.Send(recipient, message)
}
