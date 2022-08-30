package main

import (
	"fmt"
	"time"
)

func job(index int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("执行完毕 %d\n", index)
}

func main() {

	t1 := time.Now()

	for i := 0; i < 10; i++ {
		go func(index int) {
			job(index)
		}(i)

	}

	time.Sleep(5000 * time.Millisecond)

	t2 := time.Since(t1)
	fmt.Println(t2)

}
