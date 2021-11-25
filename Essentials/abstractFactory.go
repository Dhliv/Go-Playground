package Essentials

import (
	"fmt"

	"github.com/Dhliv/Go-Playground/CP"
)

type iNotificationFactory interface {
	sendNotification()
	getSender() iSender
}

type iSender interface {
	getSenderMethod() string
	getSenderChannel() string
}

type smsNotification struct {
}

func (smsNotification) sendNotification() {
	CP.Print("Sending Notification via SMS")
}

func (smsNotification) getSender() iSender {
	return smsNotificationSender{}
}

type smsNotificationSender struct {
}

func (smsNotificationSender) getSenderMethod() string {
	return "SMS"
}

func (smsNotificationSender) getSenderChannel() string {
	return "Twilio"
}

type emailNotification struct {
}

func (emailNotification) sendNotification() {
	CP.Print("Sending Notification via Email")
}

type emailNotificationSender struct {
}

func (emailNotificationSender) getSenderMethod() string {
	return "Email"
}

func (emailNotificationSender) getSenderChannel() string {
	return "SES"
}

func (emailNotification) getSender() iSender {
	return emailNotificationSender{}
}

func getNotificactionFactory(notificationType string) (iNotificationFactory, error) {
	if notificationType == "SMS" {
		return &smsNotification{}, nil
	}

	if notificationType == "Email" {
		return &emailNotification{}, nil
	}

	return nil, fmt.Errorf("No Notification Type")
}

func sendNotification(f iNotificationFactory) {
	f.sendNotification()
}

func getMethod(f iNotificationFactory) {
	CP.Print(f.getSender().getSenderMethod())
}

func AbstractFactory() {
	smsFactory, _ := getNotificactionFactory("SMS")
	emailFactory, _ := getNotificactionFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
