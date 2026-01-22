package notify

// Notifier  общий контракт для уведомлений (email/sms/мессенджер и т.д.).
type Notifier interface {
	Send(customer string, message string) error
}
