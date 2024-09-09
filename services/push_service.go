package services

import "fmt"

// PushNotification is a concrete implementation of Notification for sending push notifications.
type PushNotification struct{}

// Send sends s push notification to the recipient.
func (p *PushNotification) Send(recipient string, message string) error {
	fmt.Printf("Sending Push Notification to %s: %s\n", recipient, message)
	return nil
}
