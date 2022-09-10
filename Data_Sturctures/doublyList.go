package main

import(
	"fmt"
)

type Node struct {
	Value int
	Previous *Node
	Next *Node
}

func addNode(t *Node,v int)int{
	if root == nil {
		t = &Node{v,nil,nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:",v)
		return -1
	}

	if t.Next == nil {
		temp := t
		t.Next = &Node{v, temp, nil}
		return -2
	}
	return addNode(t.Next,v)
}


func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty List!")
		return 
	}
	for t != nil {
		fmt.Printf("%d -> ",t.Value)
		t = t.Next
	}
	fmt.Println()
}

func reverse(t *Node){
	if t == nil {
		fmt.Println("-> Empty List!")
		return
	}

	temp := t
	// 找到双链表最后一个结点，记为temp
	for t != nil {
		temp = t
		t = t.Next
	}

	// 从temp开始向前打印
	for temp.Previous != nil{
		fmt.Printf("%d -> ",temp.Value)
		temp = temp.Previous
	}

	fmt.Printf("%d -> ",temp.Value)
	fmt.Println()
}


func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty List!")
		return 0
	}
	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookupNode(t *Node, v int) bool {
	if root == nil {
		return false
	}
	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next,v)
}


var root = new(Node)

func main(){
	fmt.Println(root)
	root = nil
	traverse(root)
	addNode(root,1)
	addNode(root,1)
	traverse(root)
	addNode(root,10)
	addNode(root,5)
	addNode(root,0)
	addNode(root,0)
	addNode(root,100)
	traverse(root)
	reverse(root)
}