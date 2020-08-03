package user

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/matisszilard/devops-palinta/pkg/model"
	"github.com/sirupsen/logrus"
)

type loggingMiddleware struct {
	logger *logrus.Logger
	next   UserService
}

func (mw loggingMiddleware) GetUsers() (output []model.User, err error) {
	defer func(begin time.Time) {
		mw.logger.WithFields(logrus.Fields{
			"method": "GetUsers",
			"output": output,
			"err":    err,
			"took":   time.Since(begin),
		}).Info("Get users middleware called")
	}(time.Now())

	output, err = mw.next.GetUsers()
	return
}

type instrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           UserService
}

// NewInstrumentingMiddleware creates a new middleware for instrumenting.
func NewInstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram, countResult metrics.Histogram, next UserService) UserService {
	return instrumentingMiddleware{requestCount, requestLatency, countResult, next}
}

func (mw instrumentingMiddleware) GetUsers() (output []model.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAllDevices", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetUsers()
	return
}
