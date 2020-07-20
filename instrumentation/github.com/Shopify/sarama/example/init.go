package example

import (
	"log"

	otelglobal "go.opentelemetry.io/otel/api/global"
	oteltracestdout "go.opentelemetry.io/otel/exporters/trace/stdout"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	KafkaTopic = "sarama-instrumentation-example"
)

func InitTracer() {
	exporter, err := oteltracestdout.NewExporter(oteltracestdout.Options{PrettyPrint: true})
	if err != nil {
		log.Fatal(err)
	}
	cfg := sdktrace.Config{
		DefaultSampler: sdktrace.AlwaysSample(),
	}
	tp, err := sdktrace.NewProvider(
		sdktrace.WithConfig(cfg),
		sdktrace.WithSyncer(exporter),
	)
	if err != nil {
		log.Fatal(err)
	}
	otelglobal.SetTraceProvider(tp)
}
