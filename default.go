package caller

import (
	"context"
	"sync"
)

var (
	defaultCaller     *Caller
	defaultCallerOnce sync.Once
)

func getDefault() *Caller {
	defaultCallerOnce.Do(func() {
		defaultCaller = New()
	})
	return defaultCaller
}

func Options(ctx context.Context, url string, opts ...RequestFunc) Result {
	return getDefault().Options(ctx, url, opts...)
}

func Get(ctx context.Context, url string, opts ...RequestFunc) Result {
	return getDefault().Get(ctx, url, opts...)
}

func Head(ctx context.Context, url string, opts ...RequestFunc) Result {
	return getDefault().Head(ctx, url, opts...)
}

func Post(ctx context.Context, url string, opts ...RequestFunc) Result {
	return getDefault().Post(ctx, url, opts...)
}

func Put(ctx context.Context, url string, opts ...RequestFunc) Result {
	return getDefault().Put(ctx, url, opts...)
}

func Patch(ctx context.Context, url string, opts ...RequestFunc) Result {
	return getDefault().Patch(ctx, url, opts...)
}

func Delete(ctx context.Context, url string, opts ...RequestFunc) Result {
	return getDefault().Delete(ctx, url, opts...)
}

func Trace(ctx context.Context, url string, opts ...RequestFunc) Result {
	return getDefault().Trace(ctx, url, opts...)
}

func Connect(ctx context.Context, url string, opts ...RequestFunc) Result {
	return getDefault().Connect(ctx, url, opts...)
}
