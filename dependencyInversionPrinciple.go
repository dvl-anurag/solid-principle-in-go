package main

import "fmt"

// NotificationService is an interface representing a notification service
type NotificationService interface {
	SendNotification(message string)
}

// EmailService is a concrete implementation of NotificationService
type EmailService struct{}

// SendNotification sends a notification via email
func (e EmailService) SendNotification(message string) {
	fmt.Println("Sending email:", message)
}

// SMSService is another concrete implementation of NotificationService
type SMSService struct{}

// SendNotification sends a notification via SMS
func (s SMSService) SendNotification(message string) {
	fmt.Println("Sending SMS:", message)
}

// NotificationManager is a high-level module that depends on the NotificationService interface
type NotificationManager struct {
	Service NotificationService
}

// Notify sends a notification using the injected NotificationService
func (n NotificationManager) Notify(message string) {
	n.Service.SendNotification(message)
}
