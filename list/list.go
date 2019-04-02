package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Head *Node
}

func (l *List) addNode(val int) {
	n := &Node{Val: val}
	if l.Head == nil {
		l.Head = n
	} else {
		curr := l.Head

		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = n
	}
}
func (l *List) printList() {
	curr := l.Head

	for curr != nil {
		fmt.Println("node", curr.Val)
		curr = curr.Next
	}
}
