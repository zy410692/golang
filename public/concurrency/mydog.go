package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type MyDogStruct struct {
	funcs []MyDogFunc
	data  chan interface{}
	wg    *sync.WaitGroup
}

type MyDogFunc func() interface{}

func (this *MyDogStruct) Add(f MyDogFunc) {
	this.funcs = append(this.funcs, f)

}
func (this *MyDogStruct) do() {
	for _, f := range this.funcs {
		this.wg.Add(1)
		go func() {
			defer this.wg.Done()
			this.data <- f()
		}()
	}
}

func (this *MyDogStruct) Range(f func(value interface{})) {
	this.do()
	go func() {
		defer close(this.data)
		this.wg.Wait()
	}()
	for item := range this.data {
		f(item)

	}
}

func MyDog() *MyDogStruct {
	return &MyDogStruct{data: make(chan interface{}), wg: &sync.WaitGroup{}}
}
func main() {
	dog := MyDog()

	for i := 0; i < 10; i++ {
		dog.Add(func() interface{} {
			return rand.Intn(10)
		})

	}
	dog.Range(func(value interface{}) {
		fmt.Println(value)
	})

}
