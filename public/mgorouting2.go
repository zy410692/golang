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

	for i := 0; i < num; i++ {
		fmt.Println(job(i))
	}

	end := time.Since(start)

	fmt.Println("耗时：", end.String())

	//耗时： 3.004248843s

}
