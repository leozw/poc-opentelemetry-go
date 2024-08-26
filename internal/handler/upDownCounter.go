package handler

import (
    "log"
    "modulo/internal/observability"
    "net/http"

    "go.opentelemetry.io/otel/metric"
)

var (
    upDownCounterMeter = observability.InitMeter().Meter("updown-counter-meter")
    activeUsersCounter metric.Int64UpDownCounter
)

func init() {
    var err error
    activeUsersCounter, err = upDownCounterMeter.Int64UpDownCounter("active_users")
    if err != nil {
        log.Fatalf("Failed to create UpDownCounter: %v", err)
    }
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    activeUsersCounter.Add(r.Context(), 1)
    w.Write([]byte("User logged in"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    activeUsersCounter.Add(r.Context(), -1)
    w.Write([]byte("User logged out"))
}
