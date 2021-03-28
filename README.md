# Caller

A tool can help coder reduce many duplicate code of requesting http api, and support optional config to modify the param of
request and http client. What's more, this tool support that the response body of restful api is parsed by function and
received into the pointer of struct or map, slice.

一个工具可以帮助减少许多用于请求http api的重复代码，并支持可选的配置来修改请求和http client的参数。
此外，该工具支持接收struct、map、slice的指针来函数解析（默认是json）api的响应体。

## Usages

### Default

```go
func TestCaller(t *testing.T) {
	ctx := context.Background()
	result := map[string]interface{}{}
	err := Get(ctx, "http://127.0.0.1:8888/ping").Parse(&result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
```

### Optional

```go
func TestOptionalCaller(t *testing.T) {
    ctx := context.Background()
    result := map[string]interface{}{}
    header := map[string]string{"key": "value"}
    result := map[string]interface{}{} 
    caller := NewCaller(WithTimeout(5*time.Second), WithRetry(3, 5*time.Second))
    err := caller.Do(ctx, "http://127.0.0.1:8888/ping", WithMethod("get"), WithHeader(header)).Parse(&result)
    if err != nil { 
    	t.Fatal(err)
    }
    t.Log(result)
}
```
