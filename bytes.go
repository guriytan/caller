package caller

import (
	"bytes"
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
	if _, err := buf.ReadFrom(date); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
