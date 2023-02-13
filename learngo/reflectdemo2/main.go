package main

import (
	"fmt"
	"reflect"
)

func main() {
	type T struct {
		A int
		B string
	}

	t := T{34, "zhangyi"}
	s := reflect.ValueOf(&t).Elem()
	typeofT := s.Type()
	for i := 0; i < typeofT.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d:%s  %s:%v\n", i, typeofT.Field(i).Name, f.Type(), f.Interface())
	}
}
