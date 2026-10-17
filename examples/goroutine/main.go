package main

import (
	"fmt"
	"sync"
	"time"
)

func work(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Println(name, "执行第", i, "次")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go work("A", &wg)
	go work("B", &wg)

	wg.Wait()

	fmt.Println("所有任务执行完毕")
}
