package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/internal/entity"
	"project/pkg/logger"
	"testing"
)

func TestUserHandler_Create(t *testing.T) {
	logger := logger.NewLogger()
	mockDAO := newMockUserDAO()
	handler := NewUserHandler(mockDAO, logger)

	user := &entity.User{
		Name:  "Test User",
		Email: "test@example.com",
	}

	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.Create(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestUserHandler_GetByID(t *testing.T) {
	logger := logger.NewLogger()
	mockDAO := newMockUserDAO()
	handler := NewUserHandler(mockDAO, logger)

	req := httptest.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()

	handler.GetByID(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
