package main

import (
    "context"
    "log"
    "net/http"
    "modulo/internal/observability"
    "modulo/internal/server"
    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
    tp := observability.InitTracer()
    defer func() { _ = tp.Shutdown(context.Background()) }()

    mp := observability.InitMeter()
    defer func() { _ = mp.Shutdown(context.Background()) }()

    router := server.NewRouter()
    wrappedRouter := otelhttp.NewHandler(router, "http-server")

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", wrappedRouter))
}
