package main

import (
	"fmt"
	"sync"
)

// 问题2描述

// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
// A1B2C3D4E5F6G7H8I9J10K11L12M13N14O15P16Q17R18S19T20U21V22W23X24Y25Z26

// 思路
// 使用两个channel，一个负责打印数字，一个负责打印字母。一个channel的内容取完了通知另一个channel取字符，反之亦然。


func main() {
	// 定义channel: letter和number
	letter, number := make(chan bool), make(chan bool)
	// 定义waitGroup, 实现协程的异步
	wait := sync.WaitGroup{}

	go func ()  {
					i := 1
					for {
						select{
						case <- number:
							fmt.Print(i)
							i++
							letter <- true				
						}
					}
	}()

	wait.Add(1)
	go func(wait *sync.WaitGroup){
		i := 'A'
		for{
			select{
			case <- letter:
				if i > 'Z'{
					wait.Done()
					return
				}
				
				fmt.Print(string(i))
				i++
				
				number <- true

			}
		}
	}(&wait)

	// // 先打印数字
	// number <- true
	// wait.Wait()
	// 先打印字母
	letter <- true
	wait.Wait()

}
