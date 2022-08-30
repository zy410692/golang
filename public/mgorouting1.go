package main

import (
	"fmt"
	"sync"
	"time"
)

func job(index int) int {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("执行完毕 %d\n", index)

	return index
}

func main() {

	start := time.Now()
	num := 6
	wg := sync.WaitGroup{}

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			fmt.Println(job(index))
		}(i)

	}
	wg.Wait()
	end := time.Since(start)

	fmt.Println("耗时：", end.String())

	//耗时： 500.427361ms

}
