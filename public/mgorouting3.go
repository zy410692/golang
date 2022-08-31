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
	num := 5
	result := make(chan int)

	wg := sync.WaitGroup{}

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			result <- job(index)
		}(i)

	}

	go func() {
		defer close(result)
		wg.Wait()
	}()

	for item := range result {
		fmt.Printf("取到结果%d\n", item)
	}

	end := time.Since(start)

	fmt.Println("耗时：", end.String())

	//耗时： 500.427361ms

}
