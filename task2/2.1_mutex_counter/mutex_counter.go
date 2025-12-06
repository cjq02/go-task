package main

import (
	"fmt"
	"sync"
)

// MutexCounter 使用 sync.Mutex 保护共享计数器
func MutexCounter() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d (期望: 10000)\n", counter)
}

// NoLockCounter 不使用锁的计数器示例
// 演示在没有锁保护的情况下，并发访问共享变量会导致数据竞争和不正确的结果
func NoLockCounter() {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d (期望: 10000)\n", counter)
	fmt.Println("注意: 由于没有锁保护，实际值通常小于10000，说明发生了数据竞争")
}

func main() {
	fmt.Println("=== 使用 Mutex 的计数器 ===")
	MutexCounter()
	fmt.Println()

	fmt.Println("=== 不使用锁的计数器（演示数据竞争） ===")
	NoLockCounter()
}
