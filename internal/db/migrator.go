package db

import (
	"database/sql"
	"log"
	"project/internal/db/migrations"
)

type Migrator struct {
	db *sql.DB
}

func NewMigrator(db *sql.DB) *Migrator {
	return &Migrator{db: db}
}

func (m *Migrator) RunMigrations() error {
	log.Println("Running database migrations...")

	// マイグレーションの実行
	if err := migrations.CreateUsersTable(m.db); err != nil {
		return err
	}

	log.Println("Migrations completed successfully")
	return nil
}
