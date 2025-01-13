package orm

import (
	"context"
	"database/sql"
	"fmt"
	"project/internal/dao"
	"project/internal/entity"
	"testing"
)

func BenchmarkUserDAO_Insert(b *testing.B) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=testdb sslmode=disable")
	if err != nil {
		b.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userDAO := dao.NewUserDAO(db)
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		user := &entity.User{
			Name:  "Benchmark User",
			Email: fmt.Sprintf("benchmark%d@test.com", i),
		}
		_, err := userDAO.Insert(ctx, user)
		if err != nil {
			b.Fatalf("Insert failed: %v", err)
		}
	}
}

func BenchmarkUserDAO_FindByID(b *testing.B) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=testdb sslmode=disable")
	if err != nil {
		b.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	userDAO := dao.NewUserDAO(db)
	ctx := context.Background()

	// Create test user
	user := &entity.User{
		Name:  "Benchmark User",
		Email: "benchmark@test.com",
	}
	id, err := userDAO.Insert(ctx, user)
	if err != nil {
		b.Fatalf("Failed to create test user: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := userDAO.FindByID(ctx, id)
		if err != nil {
			b.Fatalf("FindByID failed: %v", err)
		}
	}
}
