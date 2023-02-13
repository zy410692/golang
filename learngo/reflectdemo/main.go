package main

import (
	"fmt"
	"reflect"
)

func main() {

	//简单变量一般TypeOf取到类型，对于复杂的需要用的TypeOf.Kind
	//var a int
	type A struct{ a string }
	typeOfA := reflect.TypeOf(&A{a: "111"})
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	typeofAElem := typeOfA.Elem()
	fmt.Println(typeofAElem.Name(), typeofAElem.Kind())

	type cat struct {
		Name string
		Age  int `json:"type" id:"100"`
	}
	ins := cat{Name: "qiqi", Age: 6}
	typeOfcat := reflect.TypeOf(ins)

	fmt.Println(typeOfcat.Name(), typeOfcat.Kind())

	for i := 0; i < typeOfcat.NumField(); i++ {

		field := typeOfcat.Field(i)
		fmt.Println(field.Name, field.Tag)

	}

	if catType, ok := typeOfcat.FieldByName("Age"); ok {
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}

}
