package caller

import (
	"context"
	"net/http"
	"time"
)

type Client struct {
	core Core

	retryTime     int
	retryInternal time.Duration
}

func NewClient(opts ...ConfigFunc) *Client {
	cfg := newDefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	client := &Client{
		core:          newCore(cfg),
		retryTime:     cfg.RetryTime,
		retryInternal: cfg.RetryInternal,
	}
	return client
}

func (c *Client) Do(ctx context.Context, url string, opts ...RequestFunc) Result {
	var param = &request{method: "GET"}
	for _, opt := range opts {
		opt(param)
	}
	req, err := http.NewRequestWithContext(ctx, param.method, url, param.body)
	if err != nil {
		return newErrResult(newResultError("new request failed", err))
	}
	var resp *http.Response
	do := func(ctx context.Context) error {
		var err error
		if resp, err = c.core.Do(req); err != nil {
			return err
		}
		return nil
	}
	if err = newRetry(c.retryTime, c.retryInternal).Do(ctx, do); err != nil {
		return newErrResult(err)
	}
	return newResult(resp.Body)
}

func (c *Client) Options(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("OPTIONS"))
	return c.Do(ctx, url, opts...)
}

func (c *Client) Get(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("GET"))
	return c.Do(ctx, url, opts...)
}

func (c *Client) Head(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("HEAD"))
	return c.Do(ctx, url, opts...)
}

func (c *Client) Post(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("POST"))
	return c.Do(ctx, url, opts...)
}

func (c *Client) Put(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("PUT"))
	return c.Do(ctx, url, opts...)
}

func (c *Client) Delete(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("DELETE"))
	return c.Do(ctx, url, opts...)
}

func (c *Client) Trace(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("TRACE"))
	return c.Do(ctx, url, opts...)
}

func (c *Client) Connect(ctx context.Context, url string, opts ...RequestFunc) Result {
	opts = append(opts, WithMethod("CONNECT"))
	return c.Do(ctx, url, opts...)
}
