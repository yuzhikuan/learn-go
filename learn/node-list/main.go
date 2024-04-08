package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func ShowNode(p *Node) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func main() {
	var head, node1, node2 *Node = new(Node), new(Node), new(Node)
	head.data = 1
	head.next = node1

	node1.data = 2
	node1.next = node2

	node2.data = 3

	// 头插
	var newHead *Node = new(Node)
	newHead.data = 0
	newHead.next = head

	// 尾插
	var tail *Node = new(Node)
	tail.data = 4
	node2.next = tail

	ShowNode(newHead)
}
