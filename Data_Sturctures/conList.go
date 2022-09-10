package main

import(
	"container/list"
	"fmt"
	"strconv"
)

func printList(l *list.List){
	// 从后向前打印
	for t := l.Back(); t != nil; t = t.Prev(){
		fmt.Print(t.Value," ")
	}
	fmt.Println()
  // 从前向后打印
	for t := l.Front(); t != nil; t = t.Next(){
		fmt.Print(t.Value," ")
	}
	fmt.Println()
}

func main(){
	values := list.New()
	e1 := values.PushBack("One")
	e2 := values.PushBack("Two")
	values.PushFront("Three")
	values.InsertBefore("Four",e1)
	values.InsertAfter("Five",e2)
	printList(values)
	values.Remove(e2)
	// values.Remove(e2)
	// printList(values)
	fmt.Printf("After Init() %v\n", values)

	for i := 0; i < 20; i++{
		values.PushFront(strconv.Itoa(i))
	}
	printList(values)
}