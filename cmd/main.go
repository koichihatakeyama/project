package main

import (
	"context"
	"database/sql"
	"project/internal/dao"
	"project/pkg/errors"
	"project/pkg/logger"
	"project/pkg/orm/config"

	_ "github.com/lib/pq"
)

func main() {
	logger := logger.NewLogger()
	ctx := context.Background()

	cfg := config.NewConfig()
	cfg.DriverName = "postgres"
	cfg.DSN = "host=db port=5432 user=postgres password=postgres dbname=testdb sslmode=disable"

	db, err := sql.Open(cfg.DriverName, cfg.DSN)
	if err != nil {
		logger.Error("Failed to connect to database: %v", err)
		panic(errors.NewAppError("DB_CONNECTION_ERROR", "Database connection failed", err))
	}
	defer db.Close()

	migrator := db.NewMigrator(db)
	if err := migrator.RunMigrations(); err != nil {
		logger.Error("Failed to run migrations: %v", err)
		panic(errors.NewAppError("MIGRATION_ERROR", "Database migration failed", err))
	}

	userDAO := dao.NewUserDAO(db)
	logger.Info("Application started successfully!")
}
