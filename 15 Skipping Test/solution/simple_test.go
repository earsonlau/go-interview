package main

import (
	"sort"
	"fmt"
	"testing"
)
// 使用 go test 跳过一些测试
//  Skip some tests with go test
// 是否可以跳过/排除使用go test运行的某些测试？

// 我有相当多的集成类型测试，它们调用编写为标准go测试的服务，并与go test%一起运行.当开发一个新功能时，能够跳过一些测试有时是很有用的
// ，例如，如果新功能还没有部署到测试服务器上，而我仍然想运行所有现有的测试(除了那些测试新功能的新测试)

type SumTest struct {
	testValues []int
	expectedResult int
}

func TestSum(t *testing.T) {
	testVals := []SumTest {
				{ testValues: []int{10, 20, 30}, expectedResult: 10},
				{ testValues: []int{ -10, 0, -10 }, expectedResult: -20},
				{ testValues: []int{ -10, 0, -10 }, expectedResult: -20},
				}
		for index, testVal := range testVals {
				t.Run(fmt.Sprintf("Sum #%v", index), func(subT *testing.T) {
					// 这样可以跳过
				if (t.Failed()) {
					subT.SkipNow()
					}
				_, sum := sortAndTotal(testVal.testValues)
			if (sum != testVal.expectedResult) {
			subT.Fatalf("Expected %v, Got %v", testVal.expectedResult, sum)
					}
			})
		}
}

func TestSort(t *testing.T){
	slices := [][]int {
		{ 1, 279, 48, 12, 3 },
		{ -10, 0, -10 },
		{ 1, 2, 3, 4, 5, 6, 7 },
		{ 1 },
	}
	for index, data := range slices {
		t.Run(fmt.Sprintf("Sort #%v", index), func(subT *testing.T){
			sorted, _ := sortAndTotal(data)
			if (!sort.IntsAreSorted(sorted)){
				subT.Fatalf("Unsorted data %v", sorted)
			}
		})
	}
}