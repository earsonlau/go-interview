package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func BenchmarkSort(b *testing.B){
	rand.Seed(time.Now().UnixNano())
	sizes := []int { 10, 100, 250 }
 	for _, size := range sizes {
 		b.Run(fmt.Sprintf("Array Size %v", size), func(subB *testing.B) {
 		data := make([]int, size)
 		subB.ResetTimer()
 		for i := 0; i < subB.N; i++ {
 			subB.StopTimer()
 			for j := 0; j < size; j++ {
 				data[j] = rand.Int()
 						}
 			subB.StartTimer()
 			sortAndTotal(data)
 }
 })
 }
}
	

func sortAndTotal(vals []int) (sorted []int, total int) {
	sorted = make([]int, len(vals))
	copy(sorted, vals)
	sort.Ints(sorted)
	for _, val := range sorted {
	total += val
	//total++
	}
	return
 }