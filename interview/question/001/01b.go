package main

import (
	"fmt"
	"time"
)

// 交替打印出 1122aabb3344ccdd
func main() {

	chan1 := make(chan bool)
	chan2 := make(chan bool)

	go func() {
		i := 1
		select {
		case <-chan1:
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			chan2 <- true

		}
	}()

	go func() {
		j := 'A'
		select {
		case <-chan2:
			if j > 'Z' {
				return
			}
			fmt.Print(j)
			j++
			fmt.Print(j)
			j++
			chan1 <- true
		}
	}()

	time.Sleep(1000 * time.Second)
}
