package device

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/matisszilard/devops-palinta/pkg/model"
)

type loggingMiddleware struct {
	logger log.Logger
	next   StringService
}

func (mw loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw loggingMiddleware) GetDevices() (output []model.Device, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "get all",
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.GetDevices()
	return
}

type instrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           StringService
}

// NewInstrumentingMiddleware creates a new middleware for instrumenting.
func NewInstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram, countResult metrics.Histogram, next StringService) StringService {
	return instrumentingMiddleware{requestCount, requestLatency, countResult, next}
}

func (mw instrumentingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "uppercase", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.Uppercase(s)
	return
}

func (mw instrumentingMiddleware) GetDevices() (output []model.Device, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetAllDevices", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.GetDevices()
	return
}
