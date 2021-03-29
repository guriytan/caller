package caller

import (
	"context"
	"sync"
)

var (
	defaultCaller     *Caller
	defaultCallerOnce sync.Once
)

func GetDefaultCaller() *Caller {
	defaultCallerOnce.Do(func() {
		defaultCaller = NewCaller()
	})
	return defaultCaller
}

func Options(ctx context.Context, url string, opts ...RequestFunc) Result {
	return GetDefaultCaller().Options(ctx, url, opts...)
}

func Get(ctx context.Context, url string, opts ...RequestFunc) Result {
	return GetDefaultCaller().Get(ctx, url, opts...)
}

func Head(ctx context.Context, url string, opts ...RequestFunc) Result {
	return GetDefaultCaller().Head(ctx, url, opts...)
}

func Post(ctx context.Context, url string, opts ...RequestFunc) Result {
	return GetDefaultCaller().Post(ctx, url, opts...)
}

func Put(ctx context.Context, url string, opts ...RequestFunc) Result {
	return GetDefaultCaller().Put(ctx, url, opts...)
}

func Delete(ctx context.Context, url string, opts ...RequestFunc) Result {
	return GetDefaultCaller().Delete(ctx, url, opts...)
}

func Trace(ctx context.Context, url string, opts ...RequestFunc) Result {
	return GetDefaultCaller().Trace(ctx, url, opts...)
}

func Connect(ctx context.Context, url string, opts ...RequestFunc) Result {
	return GetDefaultCaller().Connect(ctx, url, opts...)
}
