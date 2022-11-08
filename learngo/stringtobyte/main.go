package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 第一种方法:
	//var aa = "aab"
	//bb := []byte(aa)
	//fmt.Println(bb)

	var str = "aab"

	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))

	fmt.Println(strHeader.Data)
	fmt.Println(strHeader.Len)
	//fmt.Println(*(*string)(unsafe.Pointer(strHeader)))

	var strBytes []byte

	ByteHeader := (*reflect.SliceHeader)(unsafe.Pointer(&strBytes))
	ByteHeader.Data = strHeader.Data
	ByteHeader.Len = strHeader.Len
	ByteHeader.Cap = strHeader.Len

	fmt.Println(*(*[]byte)(unsafe.Pointer(ByteHeader)))

}
