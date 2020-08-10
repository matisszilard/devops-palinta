package main

import (
	"net/http"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/matisszilard/devops-palinta/service/device"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/sirupsen/logrus"
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

	log := logrus.New()

	// Elastic logger
	// client, err := elastic.NewClient(elastic.SetURL("http://okd-5mthh-worker-tb667.apps.okd.codespring.ro:30029"), elastic.SetSniff(false))
	// if err != nil {
	// 	log.Panic(err)
	// }

	// hook, err := elogrus.NewAsyncElasticHook(client, "http://okd-5mthh-worker-tb667.apps.okd.codespring.ro:30029", logrus.DebugLevel, "palinta-device")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// log.Hooks.Add(hook)

	var svc device.StringService
	svc = device.New(log)
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
	http.Handle("/api/v1/devices/uppercase", uppercaseHandler)
	http.Handle("/api/v1/devices", getDevicesHandler)
	http.Handle("/api/v1/devices/metrics", promhttp.Handler())

	log.WithFields(logrus.Fields{
		"method":   "main",
		"protocol": "HTTP",
		"port":     "8080",
	}).Info("Starting palinta device service...")

	log.Info(http.ListenAndServe(":8080", nil))
}
