package caller

import (
	"io"
	"strings"
)

type request struct {
	method string
	body   io.Reader
	header map[string]string
}

type RequestFunc func(req *request)

func WithHeader(header map[string]string) RequestFunc {
	return func(req *request) {
		req.header = header
	}
}

func WithMethod(method string) RequestFunc {
	return func(req *request) {
		req.method = strings.ToUpper(method)
	}
}

func WithBody(body io.Reader) RequestFunc {
	return func(req *request) {
		req.body = body
	}
}
