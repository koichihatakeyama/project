
package main

import (
    "context"
    "database/sql"
    "log"
    "project/pkg/orm/config"
    "project/internal/dao"
)

func main() {
    ctx := context.Background()
    
    cfg := config.NewConfig()
    cfg.DriverName = "postgres"
    cfg.DSN = "host=localhost port=5432 user=postgres password=postgres dbname=testdb sslmode=disable"
    
    db, err := sql.Open(cfg.DriverName, cfg.DSN)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    userDAO := dao.NewUserDAO(db)
    // ここにビジネスロジックを実装
}
