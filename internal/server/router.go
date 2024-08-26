package server

import (
	"modulo/internal/handler"
	"net/http" 

	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return otelhttp.NewHandler(next, "HTTP Request")
	})

	r.Get("/", handler.HomeHandler)
	r.Get("/buteco", handler.ButecoHandler)
	r.Get("/health", handler.HealthHandler)
	r.Get("/sre", handler.SREHandler)
	r.Get("/memory", handler.MemoryHandler)
	r.Get("/latency", handler.LatencyHandler)
	r.Get("/login", handler.LoginHandler)
	r.Get("/logout", handler.LogoutHandler)

	return r
}
