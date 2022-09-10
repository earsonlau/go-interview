package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}
var rwmutex = sync.RWMutex{}
var squares = map[int]int {}

func calculateSquares(max, iterations int){
	for i:=0;i<iterations;i++{
		val := rand.Intn(max)
		// 读锁
		rwmutex.RLock()
		square, ok := squares[val]
		rwmutex.RUnlock()
		if(ok) {
			fmt.Println("Cached value:",val,"=",square)
		}else{
			// 写锁
			rwmutex.Lock()
			if _,ok := squares[val];!ok{
			squares[val] = int(math.Pow(float64(val),2))
			fmt.Println("Added value:",val,squares[val])
			}
			rwmutex.Unlock()
		}
	}
	waitGroup.Done()
}
func main() {
	// 使用strings.Count()
	rand.Seed(time.Now().UnixNano())
	numRoutines := 3
	waitGroup.Add(numRoutines)
	for i:=0; i < numRoutines; i++{
		go calculateSquares(10,5)
	}
	waitGroup.Wait()
	fmt.Printf("Cached Values: %v",len(squares))
	
}
