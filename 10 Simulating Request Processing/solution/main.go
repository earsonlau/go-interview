package main

import (
	"fmt"
	// "math"
	// "math/rand"
	"sync"
	"time"
)
func processRequest(wg *sync.WaitGroup, count int){
	total := 0
	for i := 0; i < count; i++ {
		fmt.Println("processing request",total)
		total++
		time.Sleep(time.Millisecond * 250)
	}
	fmt.Println(total, "Request processed" )
	wg.Done()
}
func main() {
	// 使用strings.Count()
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	fmt.Println("Request dispatched...")
	go processRequest(&waitGroup,10)
	waitGroup.Wait()
	
	
	// rand.Seed(time.Now().UnixNano())
	// numRoutines := 3
	// waitGroup.Add(numRoutines)
	// for i:=0; i < numRoutines; i++{
	// 	go readSquares(i,10,5)
	// }
	// waitGroup.Add(1)
	// go generateSquares(10)
	// waitGroup.Wait()
	// fmt.Printf("Cached Values: %v",len(squares))
}
