package caller

import (
	"io"

	json "github.com/json-iterator/go"
)

type ParseFunc func(body io.ReadCloser, receive interface{}) error

func DefaultReceiveFunc(body io.ReadCloser, receive interface{}) error {
	defer func() { _ = body.Close() }()
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bytes, receive); err != nil {
		return err
	}
	return nil
}
