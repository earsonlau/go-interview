package main

import (
	"fmt"
	"math/rand"
	"time"
)

// nil通道的使用方式
func add(c chan int) {
	sum := 0
	t := time.NewTimer(9 * time.Second)
	for {
		select {
		case input := <- c:
			sum = sum + input
			// <-t.C语句阻塞了计数器t的c通道，其时长由 time.NewTimer()指定。
			// 不要将通道c与通道t.C混淆。时间过期后，计时器将数值发送到t.C通道中。这将触发select语句的相应分支的执行。
			// 也就是将值nil分配给通道c并输出sum变量。
		case <-t.C:
			c = nil 
			fmt.Println(sum)
		}
	}
}

// 生成随机数，持续将其发送至处于开放状态的通道
func send(c chan int){
	for {
		c <- rand.Intn(10)
	}
}

func main(){
	c := make(chan int)
	go add(c)
	go send(c)

	time.Sleep(10 * time.Second)
}