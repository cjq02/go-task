package main

import "fmt"

// ChannelBuffer 实现一个带有缓冲的通道
func ChannelBuffer() {
	ch := make(chan int, 10)
	done := make(chan bool)

	go func() {
		for i := 1; i <= 100; i++ {
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
	ChannelBuffer()
}
