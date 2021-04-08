package caller

import (
	"context"
	"net/http"
)

type Caller struct {
	core Core

	retry Retry

	parseFunc ParseFunc
}

func NewCaller(opts ...ConfigFunc) *Caller {
	cfg := newDefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	return &Caller{
		core:      newCore(cfg),
		retry:     newRetry(cfg.RetryTime, cfg.RetryInternal),
		parseFunc: cfg.ParseFunc,
	}
}

func (c *Caller) Do(ctx context.Context, url string, opts ...RequestFunc) Result {
	req, err := newRequest(ctx, "GET", url, opts...)
	if err != nil {
		return newErrResult(err)
	}
	var resp *http.Response
	do := func(ctx context.Context) error {
		if resp, err = c.core.Do(req); err != nil {
			return err
		}
		return nil
	}
	if err = c.retry.Do(ctx, do); err != nil {
		return newErrResult(err)
	}
	return newResult(resp.Body, c.parseFunc)
}

func (c *Caller) Options(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, url, append(opts, WithMethod("OPTIONS"))...)
}

func (c *Caller) Get(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, url, append(opts, WithMethod("GET"))...)
}

func (c *Caller) Head(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, url, append(opts, WithMethod("HEAD"))...)
}

func (c *Caller) Post(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, url, append(opts, WithMethod("POST"))...)
}

func (c *Caller) Put(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, url, append(opts, WithMethod("PUT"))...)
}

func (c *Caller) Delete(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, url, append(opts, WithMethod("DELETE"))...)
}

func (c *Caller) Trace(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, url, append(opts, WithMethod("TRACE"))...)
}

func (c *Caller) Connect(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, url, append(opts, WithMethod("CONNECT"))...)
}
