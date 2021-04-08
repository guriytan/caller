package caller

import (
	"reflect"
	"unsafe"
)

func bytesToString(bytes []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{Data: bytesHeader.Data, Len: bytesHeader.Len}))
}
