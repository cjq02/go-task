package main

import "fmt"

// ChannelBuffer 实现一个带有缓冲的通道
// 生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
func ChannelBuffer() {
	// 创建一个带缓冲的通道，缓冲区大小为 10
	// 带缓冲通道允许发送方在缓冲区未满时继续发送，无需立即等待接收方
	ch := make(chan int, 10)
	// 创建一个完成通道，用于通知主协程消费已完成
	done := make(chan bool)

	// 启动生产者协程：向通道中发送 100 个整数
	go func() {
		// 使用 for 循环生成并发送 1 到 100 的整数
		for i := 1; i <= 100; i++ {
			ch <- i // 将整数发送到带缓冲的通道
			// 由于通道有缓冲区，发送操作不会阻塞，直到缓冲区满
		}
		// 发送完成后关闭通道，告知消费者没有更多数据
		close(ch)
	}()

	// 启动消费者协程：从通道中接收整数并打印
	go func() {
		// 使用 for range 循环从通道接收数据，直到通道关闭
		for num := range ch {
			fmt.Printf("接收到: %d\n", num)
		}
		// 接收完成后，通知主协程
		done <- true
	}()

	// 等待消费者协程完成
	<-done
}

// 注意：如果要单独运行此文件，使用命令: go run channel_buffer.go
// 如果与 channel_communication.go 一起编译，需要注释掉其中一个 main 函数
func main() {
	ChannelBuffer()
}

