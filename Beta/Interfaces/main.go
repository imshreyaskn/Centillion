package main

import "fmt"

type Notifier interface {
	sendMessage(message string) error
}

type EmailNotifier struct {
	sender string
}

type SMSNotifier struct {
	apiProvider string
}

func (e *EmailNotifier) sendMessage(message string) error {
	fmt.Printf("Sending email from '%s' with message: '%s'\n", e.sender, message)
	return nil
}

func (s *SMSNotifier) sendMessage(message string) error {
	fmt.Printf("Sending SMS via '%s' with message: '%s'\n", s.apiProvider, message)
	return nil
}

func notify(n Notifier, message string) {
	fmt.Println("--- Notifying User ---")
	err := n.sendMessage(message)
	if err != nil {
		fmt.Printf("Error sending notification: %v\n", err)
	} else {
		fmt.Println("Notification sent successfully!")
	}
}

func main() {
	emailSender := &EmailNotifier{
		sender: "xx@gmail,com",
	}

	smsSender := &SMSNotifier{
		apiProvider: "BSNL",
	}

	notify(emailSender, "Order Confirmed")
	notify(smsSender, "Order Confirmed")
}
