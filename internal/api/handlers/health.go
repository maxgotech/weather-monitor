package handlers

import (
	"net/http"
)

func (h *Handlers) HealthHandler(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, map[string]string{
		"message": "OK",
	})
}
