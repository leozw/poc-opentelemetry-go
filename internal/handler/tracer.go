package handler

import (
    "go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("modulo-handler")
