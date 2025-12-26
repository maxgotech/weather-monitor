package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"weather-monitor/internal/models"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	user := &models.User{
		ID:       uuid.NewString(),
		Email:    req.Email,
		Password: req.Password,
	}

	models.Users[user.Email] = user
	w.WriteHeader(http.StatusCreated)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	user, ok := models.Users[req.Email]
	if !ok || user.Password != req.Password {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
}
