package order

// RepositoryInitializer — отдельная ответственность: подготовка хранилища (таблицы/схема).
// Это демонстрация Interface Segregation Principle: инициализация отдельно от записи.
type RepositoryInitializer interface {
	Init() error
}

// RepositoryWriter — контракт на сохранение заказа (не важно: sqlite, память, postgres и т.д.).
type RepositoryWriter interface {
	Save(order Order) error
}
