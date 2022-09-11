package main

import(
	"fmt"
	"time"
)

//A被阻塞，直到通道a关闭
func A(a, b chan struct{}){
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

//B被阻塞，直到通道a关闭
func B(a, b chan struct{}){
	<-a
	fmt.Println("B()!")
	close(b)
}

//C被阻塞，直到通道a关闭
func C(a chan struct{}){
	<-a
	fmt.Println("C()!")
}

func main(){
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})
	go C(z)
	go A(x,y)
	go C(z)
	go B(y,z)
	go C(z)
	close(x)
	time.Sleep(3 * time.Second)
}