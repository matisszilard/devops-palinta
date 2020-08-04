package main

import (
	"net"
	"net/http"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/matisszilard/devops-palinta/service/user"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func main() {
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "palinta",
		Subsystem: "user_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "palinta",
		Subsystem: "user_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "palinta",
		Subsystem: "user_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	logrus.SetFormatter(&logrus.JSONFormatter{})

	log := logrus.New()
	// Logstash logger
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "palinta-user"}))
	log.Hooks.Add(hook)

	var svc user.UserService
	svc = user.New(log)
	svc = user.NewInstrumentingMiddleware(requestCount, requestLatency, countResult, svc)

	// Setup routing
	http.Handle("/api/v1/users", httptransport.NewServer(
		user.MakeGetUsersEndpoint(svc),
		user.DecodeGetUsersRequest,
		user.EncodeResponse,
	))
	http.Handle("/api/v1/users/metrics", promhttp.Handler())

	log.WithFields(logrus.Fields{
		"method":   "main",
		"protocol": "HTTP",
		"port":     "8080",
	}).Info("Starting palinta user service ...")

	log.Info(http.ListenAndServe(":8080", nil))
}
