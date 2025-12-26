package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

func (h *Handlers) ListCities(w http.ResponseWriter, r *http.Request) {
	cities, err := h.db.ListCities(context.Background())
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})

		return
	}

	RespondJSON(w, http.StatusOK, cities)
}

func (h *Handlers) SaveUserCity(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Email string `json:"email"`
		City  string `json:"city"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	err := h.db.UpdateUserCity(context.Background(), req.Email, req.City)
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})

		return
	}
	RespondJSON(w, http.StatusCreated, nil)
}

func (h *Handlers) GetUserCity(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")

	userWithCity, err := h.db.GetUserWithCity(context.Background(), userEmail)
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})

		return
	}

	if userWithCity.City == nil {
		err := h.db.UpdateUserCity(context.Background(), userEmail, "Moscow")
		if err != nil {
			RespondJSON(w, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})

			return
		}

		userWithCity, err = h.db.GetUserWithCity(context.Background(), userEmail)
		if err != nil {
			RespondJSON(w, http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})

			return
		}
	}

	RespondJSON(w, http.StatusOK, map[string]string{
		"user": userWithCity.Email,
		"city": userWithCity.City.Name,
	},
	)
}
