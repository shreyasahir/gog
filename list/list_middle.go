package main

import "fmt"

func main() {
	l := &List{}
	l.addNode(1)
	l.addNode(2)
	// l.addNode(3)
	// l.addNode(4)
	// l.addNode(5)
	// l.addNode(6)
	l.printList()
	fmt.Println(findMiddle(l.Head))
}

func findMiddle(node *Node) *Node {
	if node == nil {
		return nil
	}

	slow := node
	fast := node.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
