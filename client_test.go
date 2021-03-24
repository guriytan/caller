package caller

import (
	"context"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	ctx := context.Background()
	result := map[string]interface{}{}
	err := Get(ctx, "http://127.0.0.1:8888/ping").Parse(&result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

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
