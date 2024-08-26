package handler

import (
	"log"
	"modulo/internal/observability"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/metric"
)

var (
	histogramMeter   = observability.InitMeter().Meter("histogram-meter")
	requestHistogram metric.Int64Histogram
)

func init() {
	var err error
	requestHistogram, err = histogramMeter.Int64Histogram("request_duration", metric.WithDescription("Duration of HTTP requests"))
	if err != nil {
		log.Fatalf("Failed to create Histogram: %v", err)
	}
}

func LatencyHandler(w http.ResponseWriter, r *http.Request) {
	_, span := observability.InitTracer().Tracer("modulo-handler").Start(r.Context(), "LatencyHandler")
	defer span.End()

	startTime := time.Now()
	time.Sleep(100 * time.Millisecond)
	duration := time.Since(startTime).Milliseconds()
	requestHistogram.Record(r.Context(), duration)
	w.Write([]byte("Latency recorded"))
}
