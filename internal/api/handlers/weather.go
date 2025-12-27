package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")

	coords, err := h.db.GetCityByName(context.Background(), city)
	if err != nil {
		RespondJSON(w, http.StatusNotFound, map[string]string{
			"error": "city not found",
		})

		return
	}

	weather, err := h.openMeteo.GetWeather(coords.Latitude, coords.Longitude)
	if err != nil {
		RespondJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch weather",
		})

		return
	}

	RespondJSON(w, http.StatusOK, weather)
}
