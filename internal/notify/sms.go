package notify

import "fmt"

type SMSSender struct{}

func NewSMSSender() *SMSSender {
	return &SMSSender{}
}

func (s *SMSSender) Send(customer string, message string) error {
	fmt.Printf("[SMS] клиенту %s: %s\n", customer, message)
	return nil
}
