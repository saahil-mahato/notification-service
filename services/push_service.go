package services

import "github.com/sirupsen/logrus"

// PushNotification is a concrete implementation of Notification for sending push notifications.
type PushNotification struct{}

// Send sends s push notification to the recipient.
func (p *PushNotification) Send(recipient string, message string) error {
	logrus.Infof("Sending Push Notification to %s: %s\n", recipient, message)
	return nil
}
