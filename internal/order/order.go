package order

// Order  доменная модель заказа.
// Products оставляем []string: репозиторий сам решит, как хранить (json/строка/таблица).
type Order struct {
	ID       int
	Customer string
	Products []string
	Total    float64
	Status   string
}
