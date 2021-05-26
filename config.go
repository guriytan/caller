package caller

import (
	"net/http"
	"net/url"
	"time"
)

type option struct {
	Timeout     time.Duration
	ConnTimeout time.Duration
	KeepAlive   time.Duration

	WriteBuffer int
	ReadBuffer  int
	MaxIdleConn int

	RetryTime     int
	RetryInternal time.Duration

	Proxy     func(*http.Request) (*url.URL, error)
	Redirect  func(req *http.Request, via []*http.Request) error
	CookieJar http.CookieJar

	ParseFunc ParseFunc
}

func newOption() *option {
	return &option{
		ConnTimeout: 5 * time.Second,
		Timeout:     10 * time.Second,
		KeepAlive:   10 * time.Second,
		ParseFunc:   defaultReceiveFunc,
	}
}

func (o *option) apply(opts ...Option) *option {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type Option func(config *option)

func WithTimeout(timeout time.Duration) Option {
	return func(config *option) {
		config.Timeout = timeout
	}
}

func WithConnTimeout(timeout time.Duration) Option {
	return func(config *option) {
		config.ConnTimeout = timeout
	}
}

func WithKeepAlive(alive time.Duration) Option {
	return func(config *option) {
		config.KeepAlive = alive
	}
}

func WithWriteBuffer(buffer int) Option {
	return func(config *option) {
		config.WriteBuffer = buffer
	}
}

func WithReadBuffer(buffer int) Option {
	return func(config *option) {
		config.ReadBuffer = buffer
	}
}

func WithMaxIdleConn(conn int) Option {
	return func(config *option) {
		config.MaxIdleConn = conn
	}
}

func WithRetry(retries int, internal time.Duration) Option {
	return func(config *option) {
		config.RetryTime = retries
		config.RetryInternal = internal
	}
}

func WithProxyURL(proxyURL string) Option {
	return func(config *option) {
		u, _ := url.Parse(proxyURL)
		config.Proxy = http.ProxyURL(u)
	}
}

func WithProxy(proxy func(*http.Request) (*url.URL, error)) Option {
	return func(config *option) {
		config.Proxy = proxy
	}
}

func WithRedirect(redirect func(req *http.Request, via []*http.Request) error) Option {
	return func(config *option) {
		config.Redirect = redirect
	}
}

func WithCookieJar(cookiejar http.CookieJar) Option {
	return func(config *option) {
		config.CookieJar = cookiejar
	}
}

func WithParseFunc(parseFunc ParseFunc) Option {
	return func(config *option) {
		config.ParseFunc = parseFunc
	}
}
