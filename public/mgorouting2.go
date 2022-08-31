package main

import (
	"fmt"
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
	result := make(chan int)

	for i := 0; i < num; i++ {
		go func(index int) {
			result <- job(index)
		}(i)

	}

	count := 0
	for item := range result {
		count++
		fmt.Printf("取到结果%d\n", item)
		if count == num {
			close(result)
			break
		}
	}

	end := time.Since(start)

	fmt.Println("耗时：", end.String())

	//耗时： 500.427361ms

}
