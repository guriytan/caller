package caller

import (
	"context"
	"fmt"
	json "github.com/json-iterator/go"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	ctx := context.Background()
	url, shutdown := httpServer()
	defer shutdown()
	result := map[string]interface{}{}
	err := Get(ctx, url).Parse(&result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestOptionalClient(t *testing.T) {
	ctx := context.Background()

	url, shutdown := httpServer()
	defer shutdown()

	result := map[string]interface{}{}
	config := &Config{
		Timeout:       30 * time.Second,
		ConnTimeout:   10 * time.Second,
		KeepAlive:     10 * time.Second,
		RetryTime:     3,
		RetryInternal: time.Second,
	}
	caller := NewCaller(config.Options()...)
	err := caller.Do(ctx, fmt.Sprintf("%s/ping", url), WithMethod("get"), WithHeader(map[string]string{"key": "value"})).Parse(&result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func httpServer() (url string, close func()) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request url: %s\n", r.RequestURI)
		if r.RequestURI == "/ping" {
			log.Printf("header: %+v\n", r.Header)
		}
		resp := struct {
			Slice []string
			Int   int64
		}{
			Slice: []string{"1", "2", "3"},
			Int:   1000,
		}
		marshal, _ := json.Marshal(&resp)
		_, _ = w.Write(marshal)
	}))
	log.Printf("http url: %s\n", ts.URL)
	return ts.URL, ts.Close
}
