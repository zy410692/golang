package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Id   int
	Name string
}

func main() {

	user := &User{Id: 102, Name: "zy"}
	//第一种方法
	//b, err := json.Marshal(user)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(b)
	//第二种方法
	var b []byte

	var ByteHeader = (*reflect.SliceHeader)(unsafe.Pointer(&b))
	ByteHeader.Cap = (int)(unsafe.Sizeof(User{}))
	ByteHeader.Len = (int)(unsafe.Sizeof(User{}))
	ByteHeader.Data = (uintptr)(unsafe.Pointer(user))

	fmt.Println(b)

	fmt.Println(*(*User)(unsafe.Pointer(user)))

	a := (*User)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&b)).Data))
	fmt.Println(a)

}
