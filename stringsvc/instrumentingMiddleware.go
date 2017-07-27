package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           StringService
}

func (mw instrumentingMiddleware) Uppercase(ctw context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{
			Key:   "method",
			Value: "uppercase",
		}
		errorField := metrics.Field{
			Key:   "method",
			Value: fmt.Sprintf("%v", err),
		}
		mw.requestCount.With(methodField).Add(1)
		mw.requestLatency.With(methodField).With(errorField).Observe(time.Since(begin))
	}(time.Now())

	output, err := mw.next.Uppercase(ctx, s)
}
