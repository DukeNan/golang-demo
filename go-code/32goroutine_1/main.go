package main

import (
	"fmt"
	"time"
)

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d, start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker: %d end job:%d\n", id, j)
		results <- j * 2
	}

}

func main() {
	// ch := make(chan int)
	// go recv(ch)
	// ch <- 10
	// fmt.Println("发送成功")
	fmt.Println("=======有缓冲的通道=======")
	// ch1 := make(chan int, 1)
	// ch1 <- 10
	// fmt.Println("发送成功")
	fmt.Println("=======循环取值=============")
	// ch1 := make(chan int)
	// ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		ch1 <- i
	// 	}
	// 	close(ch1)
	// }()
	// // 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	// go func() {
	// 	// for {
	// 	// 	i, ok := <-ch1
	// 	// 	if !ok {
	// 	// 		break
	// 	// 	}
	// 	// 	ch2 <- i * i
	// 	// }
	// 	// close(ch2)
	// 	for i := range ch1 {
	// 		ch2 <- i * i
	// 	}
	// 	close(ch2)
	// }()
	// 在主goroutine中从ch2中接收值打印
	// for i := range ch2 { // 通道关闭后会退出for range循环
	// 	fmt.Println(i)
	// }
	fmt.Println("=========单向通道=========")
	// ch3 := make(chan int)
	// ch4 := make(chan int)
	// go counter(ch3)
	// go squarer(ch4, ch3)
	// printer(ch4)
	fmt.Println("======goroutine池==========")
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启三个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}

}
