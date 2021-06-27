package main

import (
	"fmt"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/exporters/prometheus"
)

func initMeter(mux *http.ServeMux) error {
	exporter, err := prometheus.New(prometheus.Config{}, nil)
	if err != nil {
		return fmt.Errorf("initialize prometheus pipeline: %w", err)
	}

	if err := runtime.Start(
		runtime.WithMinimumReadMemStatsInterval(time.Second),
	); err != nil {
		return fmt.Errorf("initialize runtime metrics: %w", err)
	}

	mux.Handle("/metrics", exporter)
	return nil
}
