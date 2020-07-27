package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/matisszilard/devops-palinta/service/device"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "palinta",
		Subsystem: "device_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "palinta",
		Subsystem: "device_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "palinta",
		Subsystem: "device_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	logger := log.NewLogfmtLogger(os.Stderr)

	var svc device.StringService
	svc = device.New(logger)
	svc = device.NewInstrumentingMiddleware(requestCount, requestLatency, countResult, svc)

	uppercaseHandler := httptransport.NewServer(
		device.MakeUppercaseEndpoint(svc),
		device.DecodeUppercaseRequest,
		device.EncodeResponse,
	)

	getDevicesHandler := httptransport.NewServer(
		device.MakeGetDevicesEndpoint(svc),
		device.DecodeGetDevicesRequest,
		device.EncodeResponse,
	)

	// Setup routing
	http.Handle("/api/v1/uppercase", uppercaseHandler)
	http.Handle("/api/v1/devices", getDevicesHandler)
	http.Handle("/metrics", promhttp.Handler())

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
