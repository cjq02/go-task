package main

import "fmt"

// Channel 使用通道实现两个协程之间的通信
// 一个协程生成从1到10的整数并发送到通道，另一个协程接收并打印
func Channel() {
	// 创建一个整型通道，用于两个协程之间的通信
	ch := make(chan int)
	// 创建一个完成通道，用于通知主协程接收已完成
	done := make(chan bool)

	// 启动发送协程：生成 1 到 10 的整数并发送到通道
	go func() {
		// 使用 for 循环生成 1 到 10 的整数
		for i := 1; i <= 10; i++ {
			ch <- i // 将整数发送到通道
		}
		// 发送完成后关闭通道，告知接收方没有更多数据
		close(ch)
	}()

	// 启动接收协程：从通道接收整数并打印
	go func() {
		// 使用 for range 循环从通道接收数据，直到通道关闭
		for num := range ch {
			fmt.Printf("接收到: %d\n", num)
		}
		// 接收完成后，通知主协程
		done <- true
	}()

	// 等待接收协程完成
	<-done
}

func main() {
	Channel()
}
