package handler

import (
    "net/http"
    "modulo/internal/observability"
    "modulo/pkg/response"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    _, span := observability.InitTracer().Tracer("modulo-handler").Start(r.Context(), "HomeHandler")
    defer span.End()

    response.JSON(w, http.StatusOK, map[string]string{"message": "Welcome to the home page!"})
}
