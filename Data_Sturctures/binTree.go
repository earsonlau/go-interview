package main

import(
	"fmt"
	"math/rand"
	"time"
)

type Tree struct{
	Left *Tree
	Value int
	Right *Tree
}

func traverse(t *Tree){
	if t == nil{
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value," ")
	traverse(t.Right)
}

func create(n int) *Tree{
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++{
		temp := rand.Intn(n * 2)
		fmt.Println("The value of temp is",temp)
		t = insert(t,temp)
	}
	return t
}

func insert(t *Tree, v int) *Tree{
	// 根节点
	if t == nil {
		return &Tree{nil,v,nil}
	}
	// 树中已存在v
	if v == t.Value{
		return t
	}
	// v比根节点小，插入左子树
	if v < t.Value {
		t.Left = insert(t.Left,v)
		return t
	}
	// v比根节点大，插入右子树
	t.Right = insert(t.Right,v)
	// 返回树t
	return t 
}

func main(){
	tree := create(10)
	fmt.Println("The value of root of the tree is",tree.Value)
	traverse(tree)
	fmt.Println()
	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println("The value of root of the tree is",tree.Value)
}