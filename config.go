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

type ConfigFunc func(config *Config)

func WithTimeout(timeout time.Duration) ConfigFunc {
	return func(config *Config) {
		config.Timeout = timeout
	}
}

func WithConnTimeout(timeout time.Duration) ConfigFunc {
	return func(config *Config) {
		config.ConnTimeout = timeout
	}
}

func WithKeepAlive(alive time.Duration) ConfigFunc {
	return func(config *Config) {
		config.KeepAlive = alive
	}
}

func WithWriteBuffer(buffer int) ConfigFunc {
	return func(config *Config) {
		config.WriteBuffer = buffer
	}
}

func WithReadBuffer(buffer int) ConfigFunc {
	return func(config *Config) {
		config.ReadBuffer = buffer
	}
}

func WithMaxIdleConn(conn int) ConfigFunc {
	return func(config *Config) {
		config.MaxIdleConn = conn
	}
}

func WithRetry(retries int, internal time.Duration) ConfigFunc {
	return func(config *Config) {
		config.RetryTime = retries
		config.RetryInternal = internal
	}
}

func WithProxyURL(proxyURL string) ConfigFunc {
	return func(config *Config) {
		u, _ := url.Parse(proxyURL)
		config.Proxy = http.ProxyURL(u)
	}
}

func WithRedirect(redirect func(req *http.Request, via []*http.Request) error) ConfigFunc {
	return func(config *Config) {
		config.Redirect = redirect
	}
}

func WithCookieJar(cookiejar http.CookieJar) ConfigFunc {
	return func(config *Config) {
		config.CookieJar = cookiejar
	}
}

func WithParseFunc(parseFunc ParseFunc) ConfigFunc {
	return func(config *Config) {
		config.ParseFunc = parseFunc
	}
}
