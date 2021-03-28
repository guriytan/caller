package caller

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
)

type Core interface {
	Do(req *http.Request) (resp *http.Response, err error)
}

type core struct {
	client *http.Client
}

func newCore(cfg *Config) Core {
	return &core{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: cfg.Proxy,
				DialContext: (&net.Dialer{
					Timeout:   cfg.ConnTimeout,
					KeepAlive: cfg.KeepAlive,
				}).DialContext,
				MaxIdleConns:    cfg.MaxIdleConn,
				WriteBufferSize: cfg.WriteBuffer,
				ReadBufferSize:  cfg.ReadBuffer,
			},
			CheckRedirect: cfg.Redirect,
			Jar:           cfg.CookieJar,
			Timeout:       cfg.Timeout,
		},
	}
}

func (c core) Do(req *http.Request) (resp *http.Response, err error) {
	response, err := c.client.Do(req)
	if err != nil {
		return nil, c.handleError(err)
	}
	switch {
	case response.StatusCode >= http.StatusInternalServerError:
		err = newNoRetryError(newHttpError(response.StatusCode, fmt.Sprintf("body: %s", c.handlerResponse(response)), ErrRequest))
	case response.StatusCode >= http.StatusBadRequest:
		err = newHttpError(response.StatusCode, fmt.Sprintf("body: %s", c.handlerResponse(response)), ErrServer)
	}
	return response, nil
}

func (c *core) handleError(err error) error {
	switch err := err.(type) {
	case net.Error:
		if err.Timeout() {
			return newNoRetryError(newResultError("timeout", err))
		}
	case *url.Error:
		if err.Timeout() {
			return newNoRetryError(newResultError("timeout", err))
		}
	}
	return newResultError("http failed", err)
}

func (c *core) handlerResponse(resp *http.Response) string {
	defer func() { _ = resp.Body.Close() }()
	bytes, _ := io.ReadAll(resp.Body)
	return bytesToString(bytes)
}
