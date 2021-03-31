package caller

import (
	"context"
	"net/http"
)

type Caller struct {
	core Core

	retry Retry
}

func NewCaller(opts ...ConfigFunc) *Caller {
	cfg := newDefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return &Caller{
		core:  newCore(cfg),
		retry: newRetry(cfg.RetryTime, cfg.RetryInternal),
	}
}

func (c *Caller) Do(ctx context.Context, url string, opts ...RequestFunc) Result {
	var param = &request{method: "GET"}
	for _, opt := range opts {
		opt(param)
	}
	req, err := http.NewRequestWithContext(ctx, param.method, url, param.body)
	if err != nil {
		return newErrResult(newResultError("new request failed", err))
	}
	for key, value := range param.header {
		req.Header.Set(key, value)
	}
	var resp *http.Response
	do := func(ctx context.Context) error {
		var err error
		if resp, err = c.core.Do(req); err != nil {
			return err
		}
		return nil
	}
	if err = c.retry.Do(ctx, do); err != nil {
		return newErrResult(err)
	}
	return newResult(resp.Body)
}

func (c *Caller) Options(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("OPTIONS"))
	return c.Do(ctx, url, opts...)
}

func (c *Caller) Get(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("GET"))
	return c.Do(ctx, url, opts...)
}

func (c *Caller) Head(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("HEAD"))
	return c.Do(ctx, url, opts...)
}

func (c *Caller) Post(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("POST"))
	return c.Do(ctx, url, opts...)
}

func (c *Caller) Put(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("PUT"))
	return c.Do(ctx, url, opts...)
}

func (c *Caller) Delete(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("DELETE"))
	return c.Do(ctx, url, opts...)
}

func (c *Caller) Trace(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("TRACE"))
	return c.Do(ctx, url, opts...)
}

func (c *Caller) Connect(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("CONNECT"))
	return c.Do(ctx, url, opts...)
}
