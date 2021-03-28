package caller

import "context"

var defaultCaller *Caller

func init() {
	defaultCaller = NewCaller()
}

func Options(ctx context.Context, url string, opts ...RequestFunc) Result {
	return defaultCaller.Options(ctx, url, opts...)
}

func Get(ctx context.Context, url string, opts ...RequestFunc) Result {
	return defaultCaller.Get(ctx, url, opts...)
}

func Head(ctx context.Context, url string, opts ...RequestFunc) Result {
	return defaultCaller.Head(ctx, url, opts...)
}

func Post(ctx context.Context, url string, opts ...RequestFunc) Result {
	return defaultCaller.Post(ctx, url, opts...)
}

func Put(ctx context.Context, url string, opts ...RequestFunc) Result {
	return defaultCaller.Put(ctx, url, opts...)
}

func Delete(ctx context.Context, url string, opts ...RequestFunc) Result {
	return defaultCaller.Delete(ctx, url, opts...)
}

func Trace(ctx context.Context, url string, opts ...RequestFunc) Result {
	return defaultCaller.Trace(ctx, url, opts...)
}

func Connect(ctx context.Context, url string, opts ...RequestFunc) Result {
	return defaultCaller.Connect(ctx, url, opts...)
}
