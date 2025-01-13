package main

import (
	"context"
	"database/sql"
	"log"
	"project/internal/dao"
	"project/internal/entity"
	"project/pkg/orm/config"
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

	userDAO := dao.NewUserDAO(db)

	// サンプルユーザーを作成
	user := entity.NewUser("Test User", "test@example.com")

	// ユーザーを保存
	id, err := userDAO.Insert(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created user with ID: %d\n", id)

	// 保存したユーザーを取得
	savedUser, err := userDAO.FindByID(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found user: %+v\n", savedUser)
}
