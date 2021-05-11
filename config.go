package caller

import (
	"net/http"
	"net/url"
	"time"
)

type Config struct {
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

func newDefaultConfig() *Config {
	return &Config{
		ConnTimeout: 5 * time.Second,
		Timeout:     10 * time.Second,
		KeepAlive:   10 * time.Second,
		ParseFunc:   defaultReceiveFunc,
	}
}

func (c *Config) Options() []ConfigFunc {
	return []ConfigFunc{
		WithTimeout(c.Timeout),
		WithConnTimeout(c.ConnTimeout),
		WithKeepAlive(c.KeepAlive),
		WithWriteBuffer(c.WriteBuffer),
		WithReadBuffer(c.ReadBuffer),
		WithMaxIdleConn(c.MaxIdleConn),
		WithRetry(c.RetryTime, c.RetryInternal),
		WithProxy(c.Proxy),
		WithRedirect(c.Redirect),
		WithCookieJar(c.CookieJar),
		WithParseFunc(c.ParseFunc),
	}
}

type ConfigFunc func(config *Config)

func WithTimeout(timeout time.Duration) ConfigFunc {
	return func(config *Config) {
		if timeout > 0 {
			config.Timeout = timeout
		}
	}
}

func WithConnTimeout(timeout time.Duration) ConfigFunc {
	return func(config *Config) {
		if timeout > 0 {
			config.ConnTimeout = timeout
		}
	}
}

func WithKeepAlive(alive time.Duration) ConfigFunc {
	return func(config *Config) {
		if alive > 0 {
			config.KeepAlive = alive
		}
	}
}

func WithWriteBuffer(buffer int) ConfigFunc {
	return func(config *Config) {
		if buffer > 0 {
			config.WriteBuffer = buffer
		}
	}
}

func WithReadBuffer(buffer int) ConfigFunc {
	return func(config *Config) {
		if buffer > 0 {
			config.ReadBuffer = buffer
		}
	}
}

func WithMaxIdleConn(conn int) ConfigFunc {
	return func(config *Config) {
		if conn > 0 {
			config.MaxIdleConn = conn
		}
	}
}

func WithRetry(retries int, internal time.Duration) ConfigFunc {
	return func(config *Config) {
		if retries > 0 {
			config.RetryTime = retries
			config.RetryInternal = internal
		}
	}
}

func WithProxyURL(proxyURL string) ConfigFunc {
	return func(config *Config) {
		if len(proxyURL) != 0 {
			u, _ := url.Parse(proxyURL)
			config.Proxy = http.ProxyURL(u)
		}
	}
}

func WithProxy(proxy func(*http.Request) (*url.URL, error)) ConfigFunc {
	return func(config *Config) {
		if proxy != nil {
			config.Proxy = proxy
		}
	}
}

func WithRedirect(redirect func(req *http.Request, via []*http.Request) error) ConfigFunc {
	return func(config *Config) {
		if redirect != nil {
			config.Redirect = redirect
		}
	}
}

func WithCookieJar(cookiejar http.CookieJar) ConfigFunc {
	return func(config *Config) {
		if cookiejar != nil {
			config.CookieJar = cookiejar
		}
	}
}

func WithParseFunc(parseFunc ParseFunc) ConfigFunc {
	return func(config *Config) {
		if parseFunc != nil {
			config.ParseFunc = parseFunc
		}
	}
}
