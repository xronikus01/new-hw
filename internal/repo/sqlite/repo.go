package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	"example/solid/internal/order"
)

type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{db: db}
}

// Init реализует order.RepositoryInitializer
func (r *SQLiteRepo) Init() error {
	_, err := r.db.Exec(`
CREATE TABLE IF NOT EXISTS orders (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	customer TEXT NOT NULL,
	products TEXT NOT NULL,
	total REAL NOT NULL,
	status TEXT NOT NULL
)`)
	return err
}

// Save реализует order.RepositoryWriter
func (r *SQLiteRepo) Save(o order.Order) error {
	products := strings.Join(o.Products, ",")
	_, err := r.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		o.Customer, products, o.Total, o.Status,
	)
	if err != nil {
		return fmt.Errorf("insert order: %w", err)
	}
	return nil
}
