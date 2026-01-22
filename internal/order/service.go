package order

import "example/solid/internal/notify"

type Service struct {
	writer   RepositoryWriter
	notifier notify.Notifier
}

func NewService(writer RepositoryWriter, notifier notify.Notifier) *Service {
	return &Service{
		writer:   writer,
		notifier: notifier,
	}
}

func (s *Service) CreateOrder(customer string, products []string, total float64) error {
	o := Order{
		Customer: customer,
		Products: products,
		Total:    total,
		Status:   "pending",
	}

	if err := s.writer.Save(o); err != nil {
		return err
	}

	return s.notifier.Send(customer, "Заказ создан")
}
