package caller

import (
	"context"
	"testing"
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
