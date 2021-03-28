package caller

import (
	json "github.com/json-iterator/go"
)

type ParseFunc func(data []byte, receive interface{}) error

func defaultReceiveFunc(bytes []byte, receive interface{}) error {
	if err := json.Unmarshal(bytes, receive); err != nil {
		return err
	}
	return nil
}
