package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"example/solid/internal/notify"
	"example/solid/internal/order"
	"example/solid/internal/repo/sqlite"
)

func main() {
	db, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := sqlite.NewSQLiteRepo(db)

	// отдельная ответственность: инициализация хранилища (таблица)
	if err := repo.Init(); err != nil {
		log.Fatal(err)
	}

	// 1) Сервис с Email-уведомлениями
	emailSvc := order.NewService(repo, notify.NewEmailSender())
	if err := emailSvc.CreateOrder("Иван", []string{"apple", "banana"}, 10.5); err != nil {
		log.Fatal(err)
	}

	// 2) Сервис с SMS-уведомлениями (код сервиса не меняем!)
	smsSvc := order.NewService(repo, notify.NewSMSSender())
	if err := smsSvc.CreateOrder("Пётр", []string{"milk"}, 3.2); err != nil {
		log.Fatal(err)
	}
}
