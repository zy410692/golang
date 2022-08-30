package main

import (
	"fmt"
	"time"
)

func job() chan string {
	ret := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ret <- "success"
	}()
	return ret
}

func run() (interface{}, error) {

	c := job()
	select {
	case r := <-c:
		return r, nil
	case <-time.After(time.Duration(3 * time.Second)):
		return nil, fmt.Errorf("time out ")

	}
}

func main() {
	fmt.Println(run())
}
