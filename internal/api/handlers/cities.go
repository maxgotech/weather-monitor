package handlers

import (
	"encoding/json"
	"net/http"

	"weather-monitor/internal/models"
)

func (h *Handlers) ListCities(w http.ResponseWriter, r *http.Request) {
	cities := make([]models.City, 0, len(models.AvailableCities))

	for _, city := range models.AvailableCities {
		cities = append(cities, city)
	}

	RespondJSON(w, http.StatusOK, cities)
}

func (h *Handlers) SaveUserCity(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Email string `json:"email"`
		City  string `json:"city"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	models.Users[req.Email].City = req.City
	RespondJSON(w, http.StatusCreated, nil)
}

func (h *Handlers) GetUserCity(w http.ResponseWriter, r *http.Request) {
	userEmail := r.URL.Query().Get("email")
	entry, exists := models.Users[userEmail]
	if !exists {
		RespondJSON(w, http.StatusNotFound, map[string]string{
			"error": "user not found",
		})

		return
	}

	if entry.City == "" {
		entry.City = "Moscow"
	}

	RespondJSON(w, http.StatusOK, map[string]string{
		"user": entry.Email,
		"city": entry.City,
	},
	)
}
