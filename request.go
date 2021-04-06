package caller

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type request struct {
	method string
	rawURL string
	query  url.Values
	body   io.Reader
	header map[string]string
}

func newRequest(ctx context.Context, method, url string, opts ...RequestFunc) (*http.Request, error) {
	param := &request{method: method, rawURL: url}
	for _, opt := range opts {
		opt(param)
	}
	parseURL, err := param.url()
	if err != nil {
		return nil, newResultError("parse url failed", err)
	}
	req, err := http.NewRequestWithContext(ctx, param.method, parseURL, param.body)
	if err != nil {
		return nil, newResultError("new request failed", err)
	}
	for key, value := range param.header {
		req.Header.Add(key, value)
	}
	return req, nil
}

func (r *request) url() (string, error) {
	parseURL, err := url.Parse(r.rawURL)
	if err != nil {
		return "", err
	}
	if len(parseURL.Scheme) == 0 {
		parseURL.Scheme = "http"
	}
	parseURL.RawQuery = r.query.Encode()

	return parseURL.String(), nil
}

type RequestFunc func(req *request)

func WithMethod(method string) RequestFunc {
	return func(req *request) {
		req.method = strings.ToUpper(method)
	}
}

func WithQuery(query map[string]string) RequestFunc {
	return func(req *request) {
		req.query = make(url.Values, len(query))
		for key, value := range query {
			req.query.Add(key, value)
		}
	}
}

func WithBody(body io.Reader) RequestFunc {
	return func(req *request) {
		req.body = body
	}
}

func WithHeader(header map[string]string) RequestFunc {
	return func(req *request) {
		req.header = header
	}
}
