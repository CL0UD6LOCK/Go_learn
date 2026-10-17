package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for worker := 0; worker < 2; worker++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for i := 0; i < 100000; i++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("预期结果:", 200000)
	fmt.Println("实际结果:", counter)
}
