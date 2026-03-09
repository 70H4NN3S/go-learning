package main

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Head *Node
	size int
}

func (ll *List) Push(val int) {
	node := Node{val, nil}
	if ll.size != 0 {
		node.Next = ll.Head
	}
	ll.Head = &node
}
