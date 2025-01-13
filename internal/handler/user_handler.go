package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "project/internal/dao"
    "project/internal/entity"
    "project/pkg/logger"
    "github.com/gorilla/mux"
)

type UserHandler struct {
    userDAO *dao.UserDAO
    logger  *logger.Logger
}
  func NewUserHandler(userDAO *dao.UserDAO, logger *logger.Logger) *UserHandler {
      return &UserHandler{
          userDAO:   userDAO,
          logger:    logger,
          validator: validator.NewUserValidator(),
      }
  }

  // レスポンス用の構造体を追加
  type CreateUserResponse struct {
      ID int64 `json:"id"`
  }

  func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
      var req model.CreateUserRequest
      if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
          h.logger.Error("リクエストのデコードに失敗: %v", err)
          http.Error(w, "Invalid request body", http.StatusBadRequest)
          return
      }

      if errors := h.validator.ValidateCreate(&req); len(errors) > 0 {
          w.Header().Set("Content-Type", "application/json")
          w.WriteHeader(http.StatusBadRequest)
          json.NewEncoder(w).Encode(errors)
          return
      }

      id, err := h.userDAO.Insert(r.Context(), &user)
      if err != nil {
          h.logger.Error("Failed to create user: %v", err)
          http.Error(w, err.Error(), http.StatusInternalServerError)
          return
      }

      response := CreateUserResponse{ID: id}
      w.Header().Set("Content-Type", "application/json")
      w.WriteHeader(http.StatusCreated)
      json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseInt(vars["id"], 10, 64)
    if err != nil {
        h.logger.Error("Invalid ID format: %v", err)
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    user, err := h.userDAO.FindByID(r.Context(), id)
    if err != nil {
        h.logger.Error("Failed to find user: %v", err)
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}
