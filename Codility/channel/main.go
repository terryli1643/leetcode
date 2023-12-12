package main

import (
	"fmt"
	"sync"
)

// 生产者
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		ch <- i // 将数据发送到 channel 中
	}

	// 以便消费者可以在读取完所有数据后停止阻塞并退出
	close(ch) // 关闭 channel
}

// 消费者
func consumer(ch <-chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	// 从 channel 中读取数据，直到 channel 关闭
	for data := range ch {
		fmt.Printf("consumer %d got data: %d\n", id, data)
	}
}

func main() {
	// 创建一个 channel
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(6)

	go producer(ch, &wg) // 启动生产者 goroutine

	// 启动5个消费者 goroutine
	go consumer(ch, &wg, 1)
	go consumer(ch, &wg, 2)
	go consumer(ch, &wg, 3)
	go consumer(ch, &wg, 4)
	go consumer(ch, &wg, 5)

	// 使用 sync.WaitGroup 等待所有 goroutine 执行完毕
	wg.Wait() // 等待所有 goroutine 结束
}
