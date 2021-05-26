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

func New(opts ...Option) *Caller {
	cfg := newOption().apply(opts...)
	return &Caller{
		core:      newCore(cfg),
		retry:     newRetry(cfg.RetryTime, cfg.RetryInternal),
		parseFunc: cfg.ParseFunc,
	}
}

func (c *Caller) Do(ctx context.Context, method, url string, opts ...RequestFunc) Result {
	req, err := newRequest(ctx, method, url, opts...)
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
	return c.Do(ctx, http.MethodOptions, url, opts...)
}

func (c *Caller) Get(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, http.MethodGet, url, opts...)
}

func (c *Caller) Head(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, http.MethodHead, url, opts...)
}

func (c *Caller) Post(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, http.MethodPost, url, opts...)
}

func (c *Caller) Put(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, http.MethodPut, url, opts...)
}

func (c *Caller) Patch(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, http.MethodPatch, url, opts...)
}

func (c *Caller) Delete(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, http.MethodDelete, url, opts...)
}

func (c *Caller) Trace(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, http.MethodTrace, url, opts...)
}

func (c *Caller) Connect(ctx context.Context, url string, opts ...RequestFunc) Result {
	return c.Do(ctx, http.MethodConnect, url, opts...)
}
