package services

import "fmt"

// EmailNotification is a concrete implementation of Notification for sending emails.
type EmailNotification struct{}

// Send sends an email to the recipient.
func (e *EmailNotification) Send(recipient string, message string) error {
	fmt.Printf("Sending Email to %s: %s\n", recipient, message)
	return nil
}
