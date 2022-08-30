package main

import (
	"fmt"
	"sync"
	"time"
)

func job(index int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("执行完毕 %d\n", index)
}

func main() {

	t1 := time.Now()

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			job(index)
		}(i)

		wg.Wait()

	}

	t2 := time.Since(t1)
	fmt.Println(t2)

}
