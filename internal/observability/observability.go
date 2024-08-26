package observability

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
)

func InitTracer() *trace.TracerProvider {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "localhost:4317", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create gRPC connection to collector: %v", err)
	}

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		log.Fatalf("failed to create OTLP trace exporter: %v", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("go-service"),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp
}

func InitMeter() *metric.MeterProvider {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "localhost:4317", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to create gRPC connection to collector: %v", err)
	}

	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		log.Fatalf("failed to create OTLP metric exporter: %v", err)
	}

	mp := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(metricExporter)),
		metric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("go-service"),
		)),
	)
	otel.SetMeterProvider(mp)
	return mp
}
