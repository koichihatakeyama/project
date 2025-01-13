package dao

import (
	"context"
	"database/sql"
	"project/internal/entity"
)

type UserDAO struct {
	db *sql.DB
}

func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (d *UserDAO) Insert(ctx context.Context, user *entity.User) (int64, error) {
	query := `
        INSERT INTO users (name, email, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	var id int64
	err := d.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&id)

	return id, err
}

func (d *UserDAO) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	query := `
        SELECT id, name, email, created_at, updated_at
        FROM users
        WHERE id = $1
    `

	user := &entity.User{}
	err := d.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}
