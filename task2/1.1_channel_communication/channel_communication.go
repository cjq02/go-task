package main

import "fmt"

// Channel 使用通道实现两个协程之间的通信
func Channel() {
	ch := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		for num := range ch {
			fmt.Printf("接收到: %d\n", num)
		}
		done <- true
	}()

	<-done
}

func main() {
	Channel()
}
