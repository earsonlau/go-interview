package main

import (
	"context"
	"fmt"

	// "math"
	// "math/rand"
	"sync"
	"time"
)
func processRequest(ctx context.Context,wg *sync.WaitGroup, count int){
	total := 0
	for i := 0; i < count; i++ {
		select {
		case <- ctx.Done():
			fmt.Println("Stopping processing - request cancelled")
			goto end
		default:
			fmt.Println("processing request",total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
}
	fmt.Println(total, "Request processed" )
	end:
	wg.Done()
}
func main() {
	// 使用strings.Count()
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	fmt.Println("Request dispatched...")
	ctx,cancel := context.WithCancel(context.Background())
	
	go processRequest(ctx,&waitGroup,10)
	time.Sleep(time.Second * 2)
	fmt.Println("canceling request")
	cancel()
	waitGroup.Wait()
	// fmt.Printf("Cached Values: %v",len(squares))
}
