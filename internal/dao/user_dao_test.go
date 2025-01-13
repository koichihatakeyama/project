package dao

import (
	"context"
	"database/sql"
	"project/internal/entity"
	"testing"

	_ "github.com/lib/pq"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=testdb sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	return db
}

func TestUserDAO_Insert(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	dao := NewUserDAO(db)
	ctx := context.Background()

	user := &entity.User{
		Name:  "Test User",
		Email: "test@example.com",
	}

	id, err := dao.Insert(ctx, user)
	if err != nil {
		t.Errorf("Failed to insert user: %v", err)
	}
	if id <= 0 {
		t.Error("Expected positive ID after insert")
	}
}

func TestUserDAO_FindByID(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	dao := NewUserDAO(db)
	ctx := context.Background()

	// テストデータの作成
	user := &entity.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	id, err := dao.Insert(ctx, user)
	if err != nil {
		t.Fatalf("Failed to create test data: %v", err)
	}

	// テスト実行
	found, err := dao.FindByID(ctx, id)
	if err != nil {
		t.Errorf("Failed to find user: %v", err)
	}
	if found.ID != id {
		t.Errorf("Expected ID %d, got %d", id, found.ID)
	}
	if found.Name != user.Name {
		t.Errorf("Expected name %s, got %s", user.Name, found.Name)
	}
}
