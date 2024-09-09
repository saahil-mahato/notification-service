package services

import (
	"github.com/sirupsen/logrus"
)

// SMSNotification is a concrete implementation of Notification for sending SMS.
type SMSNotification struct{}

// Send sends an SMS to the recipient.
func (s *SMSNotification) Send(recipient string, message string) error {
	logrus.Infof("Sending SMS to %s: %s\n", recipient, message)
	return nil
}
