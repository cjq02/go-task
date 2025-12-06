package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// AtomicCounter 使用原子操作实现无锁计数器
func AtomicCounter() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d (期望: 10000)\n", counter)
}

func main() {
	AtomicCounter()
}

