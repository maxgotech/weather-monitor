package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"weather-monitor/internal/models"
)

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err := h.db.CreateUser(context.Background(), user)
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})

		return
	}

	RespondJSON(w, http.StatusCreated, nil)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	user, err := h.db.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		RespondJSON(w, http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})

		return
	}

	if user.Password != req.Password {
		RespondJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "invalid credentials",
		})

		return
	}

	RespondJSON(w, http.StatusOK, map[string]string{
		"message": "access granted",
	})
}
