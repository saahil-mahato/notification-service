package services

import "github.com/sirupsen/logrus"

// EmailNotification is a concrete implementation of Notification for sending emails.
type EmailNotification struct{}

// Send sends an email to the recipient.
func (e *EmailNotification) Send(recipient string, message string) error {
	logrus.Infof("Sending Email to %s: %s", recipient, message)
	return nil
}
