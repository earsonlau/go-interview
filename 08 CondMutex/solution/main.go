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
var readyCond = sync.NewCond(rwmutex.RLocker())

var squares = map[int]int {}

func generateSquares(max int){
	rwmutex.Lock()
	fmt.Println("generating data...")
	for val:=0;val<max;val++{
		squares[val] = int(math.Pow(float64(val),2))
		
	}
	rwmutex.Unlock()
	fmt.Println("Broadcasting condition")
	readyCond.Signal()
	waitGroup.Done()
}

func readSquares(id, max, iterations int){
	readyCond.L.Lock()
	for len(squares) == 0 {
		readyCond.Wait()
	}
	for i:=0;i<iterations;i++{
		key := rand.Intn(max)
		fmt.Println("#",id," Read value",key,squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	readyCond.L.Unlock()
	waitGroup.Done()
}

func main() {
	// 使用strings.Count()
	rand.Seed(time.Now().UnixNano())
	numRoutines := 3
	waitGroup.Add(numRoutines)
	for i:=0; i < numRoutines; i++{
		go readSquares(i,10,5)
	}
	waitGroup.Add(1)
	go generateSquares(10)
	waitGroup.Wait()
	fmt.Printf("Cached Values: %v",len(squares))
}
