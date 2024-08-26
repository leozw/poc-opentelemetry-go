package handler

import (
    "net/http"
    "modulo/pkg/response"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
    response.JSON(w, http.StatusOK, map[string]string{"status": "healthy"})
}
