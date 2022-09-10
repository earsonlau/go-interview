package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func BenchmarkSort(b *testing.B){
	rand.Seed(time.Now().UnixNano())
	size := 250
	data := make([]int, size)

	b.ResetTimer()
	for i := 0; i < b.N; i++{
		b.StopTimer()
		for j := 0; j < size; j++ {
			data[j] = rand.Int()
		}
		b.StartTimer()
		sortAndTotal(data)
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