package caller

import (
	"context"
	"errors"
	"time"
)

type retry struct {
	retry    int
	internal time.Duration
}

func newRetry(retries int, internal time.Duration) *retry {
	return &retry{retry: retries, internal: internal}
}

func (r *retry) Do(ctx context.Context, f func(ctx context.Context) error) error {
	retryTime := 0
	for {
		errC := f(ctx)
		if errC == nil {
			return nil
		}
		var e NoRetryError
		if errors.As(errC, &e) {
			return e
		}
		if retryTime > r.retry {
			return newRetryError(retryTime, errC)
		}
		select {
		case <-time.After(r.internal):
			retryTime++
		}
	}
}
