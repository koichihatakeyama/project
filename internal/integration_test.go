package internal

import (
	"context"
	"database/sql"
	"project/internal/dao"
	"project/internal/entity"
	"project/pkg/orm/transaction"
	"testing"
)

func setupIntegrationTest(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=testdb sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	return db, func() {
		db.Exec("DELETE FROM users")
		db.Close()
	}
}

func TestIntegrationUserFlow(t *testing.T) {
	db, cleanup := setupIntegrationTest(t)
	defer cleanup()

	ctx := context.Background()
	userDAO := dao.NewUserDAO(db)
	tm := transaction.NewTransactionManager(db)

	t.Run("Complete user flow", func(t *testing.T) {
		err := tm.RunInTransaction(ctx, func(tx *sql.Tx) error {
			// Create user
			user := &entity.User{
				Name:  "Integration Test User",
				Email: "integration@test.com",
			}

			id, err := userDAO.Insert(ctx, user)
			if err != nil {
				return err
			}

			// Verify user creation
			found, err := userDAO.FindByID(ctx, id)
			if err != nil {
				return err
			}

			if found.Name != user.Name {
				t.Errorf("Expected name %s, got %s", user.Name, found.Name)
			}

			return nil
		})

		if err != nil {
			t.Errorf("Transaction failed: %v", err)
		}
	})
}
