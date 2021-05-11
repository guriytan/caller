package caller

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"unsafe"
)

func bytesToString(bytes []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{Data: bytesHeader.Data, Len: bytesHeader.Len}))
}

func readerToBytes(date io.Reader) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, 512))
	read, err := buf.ReadFrom(date)
	if err != nil {
		return nil, err
	}
	if read == 0 {
		return nil, errors.New("body size is zero")
	}
	return buf.Bytes(), nil
}
