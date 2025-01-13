package main

import (
    "context"
    "database/sql"
    "project/internal/dao"
    "project/internal/db"
    "project/internal/handler"
    "project/internal/server"
    "project/pkg/orm/config"
    "project/pkg/logger"
    "project/pkg/errors"
    _ "github.com/lib/pq"
)

func main() {
    logger := logger.NewLogger()
    ctx := context.Background()

    // データベース接続設定
    cfg := config.NewConfig()
    cfg.DriverName = "postgres"
    cfg.DSN = "host=db port=5432 user=postgres password=postgres dbname=testdb sslmode=disable"

    db, err := sql.Open(cfg.DriverName, cfg.DSN)
    if err != nil {
        logger.Error("Failed to connect to database: %v", err)
        panic(errors.NewAppError("DB_CONNECTION_ERROR", "Database connection failed", err))
    }
    defer db.Close()

    // マイグレーション実行
    migrator := db.NewMigrator(db)
    if err := migrator.RunMigrations(); err != nil {
        logger.Error("Failed to run migrations: %v", err)
        panic(errors.NewAppError("MIGRATION_ERROR", "Database migration failed", err))
    }

    // ハンドラーとサーバーの設定
    userDAO := dao.NewUserDAO(db)
    userHandler := handler.NewUserHandler(userDAO, logger)
    srv := server.NewServer(logger)
    srv.SetupRoutes(userHandler)

    // サーバー起動
    logger.Info("Starting application...")
    if err := srv.Start(":8080"); err != nil {
        logger.Error("Server failed to start: %v", err)
        panic(err)
    }
}
