package handler

import (
	"modulo/internal/observability"
	"net/http"
	"runtime"

	"go.opentelemetry.io/otel/metric"
)

var (
	gaugeMeter     = observability.InitMeter().Meter("gauge-meter")
	memoryGauge, _ = gaugeMeter.Int64ObservableGauge("memory_usage", metric.WithDescription("Current memory usage in bytes"))
)

func MemoryHandler(w http.ResponseWriter, r *http.Request) {
	_, span := observability.InitTracer().Tracer("modulo-handler").Start(r.Context(), "MemoryHandler")
	defer span.End()

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	w.Write([]byte("Memory usage recorded"))
}
