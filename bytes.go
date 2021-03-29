package caller

import (
	"reflect"
	"unsafe"
)

func bytesToString(bytes []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	strHeader := reflect.StringHeader{Data: bytesHeader.Data, Len: bytesHeader.Len}
	return *(*string)(unsafe.Pointer(&strHeader))
}
