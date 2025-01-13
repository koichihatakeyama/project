package main

import (
	"context"
	"database/sql"
	"log"
	"project/internal/dao"

	"project/pkg/orm/config"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	cfg := config.NewConfig()
	cfg.DriverName = "postgres"
	cfg.DSN = "host=db port=5432 user=postgres password=postgres dbname=testdb sslmode=disable"

	db, err := sql.Open(cfg.DriverName, cfg.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// マイグレーションの実行
	migrator := db.NewMigrator(db)
	if err := migrator.RunMigrations(); err != nil {
		log.Fatal(err)
	}

	userDAO := dao.NewUserDAO(db)
	log.Println("Application started successfully!")
}
