package handler

import (
	"modulo/internal/observability"
	"modulo/pkg/response"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	meter          = observability.InitMeter().Meter("modulo-meter")
	requestCounter metric.Int64Counter
)

func init() {
	var err error
	requestCounter, err = meter.Int64Counter("modulo_requests_total", metric.WithDescription("Total number of requests"))
	if err != nil {
		panic(err)
	}
}

func ButecoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := observability.InitTracer().Tracer("modulo-handler").Start(r.Context(), "ButecoHandler")
	defer span.End()

	requestCounter.Add(ctx, 1, metric.WithAttributes(attribute.String("rota", "/buteco")))

	response.JSON(w, http.StatusOK, map[string]string{"message": "Welcome to the Buteco page!"})
}
