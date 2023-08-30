package main

import (
	"fmt"
)

// 打印出1122aabb....
func main() {
	chan1 := make(chan bool)

	chan2 := make(chan bool)
	chan3 := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-chan1:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				chan2 <- true
			}
		}
	}()
	go func() {
		j := 'A'
		for {

			select {
			case <-chan2:

				fmt.Print(string(j))
				j++
				fmt.Print(string(j))
				j++
				if j > 'Z' {
					chan3 <- true
					return
				}
				chan1 <- true
			}
		}
	}()

	chan1 <- true

	<-chan3

}
