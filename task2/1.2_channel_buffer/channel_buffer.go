package main

import (
	"fmt"
	"time"
)

// ChannelBuffer 实现一个带有缓冲的通道
func ChannelBuffer() {
	ch := make(chan int, 10)
	done := make(chan bool)

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
			if i <= 10 {
				fmt.Printf("发送: %d (缓冲区中: %d 个元素)\n", i, len(ch))
			} else if i > 11 {
				fmt.Printf("发送: %d (缓冲区已满，等待消费者消费)\n", i)
			}
		}
		fmt.Println("生产者发送完成，关闭通道")
		close(ch)
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		count := 0
		for num := range ch {
			count++
			remaining := len(ch)
			fmt.Printf("接收到: %d (剩余: %d 个)\n", num, remaining)
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Printf("消费者接收完成，共接收 %d 个数据\n", count)
		done <- true
	}()

	<-done
}

func main() {
	ChannelBuffer()
}
