
package dao

import (
    "context"
    "database/sql"
)

type UserDAO struct {
    db *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO {
    return &UserDAO{db: db}
}

func (d *UserDAO) Insert(ctx context.Context, user *User) (int64, error) {
    // 実装
}
