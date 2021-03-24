# caller

A tool can help coder reduce many duplicate code to request http api, and support optional config to modify the param of
request and http client. What's more, this tool support that the response body of restful api is parsed by function and
received into the pointer of struct or map, slice.

## Usages

### Default

```go
func TestClient(t *testing.T) {
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
func TestOptionalClient(t *testing.T) {
    ctx := context.Background()
    result := map[string]interface{}{}
    client := NewClient(WithTimeout(5*time.Second), WithRetry(3, 5*time.Second))
    err := client.Do(ctx, "http://127.0.0.1:8888/ping", WithMethod("get"), WithHeader(map[string]string{"key": "value"})).Parse(&result)
    if err != nil { 
    	t.Fatal(err)
    }
    t.Log(result)
}
```
