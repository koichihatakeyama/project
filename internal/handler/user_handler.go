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
        userDAO: userDAO,
        logger:  logger,
    }
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
    var user entity.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        h.logger.Error("Failed to decode request body: %v", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id, err := h.userDAO.Insert(r.Context(), &user)
    if err != nil {
        h.logger.Error("Failed to create user: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]int64{"id": id})
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
