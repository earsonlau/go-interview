package main

import (
	"context"
	"fmt"

	// "math"
	// "math/rand"
	"sync"
	"time"
)


const (
	countKey = iota
	sleepPeriodKey
)

func processRequest(ctx context.Context,wg *sync.WaitGroup){
	total := 0
	count := ctx.Value(countKey).(int)
	sleepPeriod := ctx.Value(sleepPeriodKey).(time.Duration)
	for i := 0; i < count; i++ {
		select {
		case <- ctx.Done():
			if (ctx.Err() == context.Canceled){
			fmt.Println("Stopping processing - request cancelled")
			} else {
				fmt.Println("Stopping processing - deadline reached")
			}
			goto end

		default:
			fmt.Println("processing request",total)
			total++
			time.Sleep(sleepPeriod)
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
	ctx,_ := context.WithTimeout(context.Background(),time.Second * 2)	
	ctx = context.WithValue(ctx, countKey, 4)
	ctx = context.WithValue(ctx, sleepPeriodKey, time.Millisecond * 250)
	go processRequest(ctx,&waitGroup)
	waitGroup.Wait()
	}
