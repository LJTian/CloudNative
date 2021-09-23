package main

import "fmt"

// 生产者
func WriteFunc(ch1 chan int, intNum int) {

	ch1 <- intNum

}

// 消费者
func ReadFunc(ch1 chan int, okCh chan int) {
	i := <-ch1
	fmt.Println("读出的内容为：", i, "\n")
	okCh <- 1
}

// GoroutineChannel 携程及管道，有缓冲
func GoroutineChannelMain() {
	ch1 := make(chan int, 10)
	OkCh := make(chan int, 10)

	for i := 0; i < 10; i++ {
		go WriteFunc(ch1, i)
	}
	for i := 0; i < 10; i++ {
		go ReadFunc(ch1, OkCh)
	}

	for i := 0; i < 10; i++ {
		<-OkCh
	}
	fmt.Println("结束")
}
