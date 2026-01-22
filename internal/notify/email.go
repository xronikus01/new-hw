package notify

import "fmt"

type EmailSender struct{}

func NewEmailSender() *EmailSender {
	return &EmailSender{}
}

func (s *EmailSender) Send(customer string, message string) error {
	fmt.Printf("[EMAIL] клиенту %s: %s\n", customer, message)
	return nil
}
