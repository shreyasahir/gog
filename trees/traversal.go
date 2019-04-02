package main

import "fmt"

func main() {
	t := &tree{}
	t.addNode(8)
	t.addNode(6)
	t.addNode(5)
	t.addNode(15)
	t.addNode(29)
	t.addNode(44)
	t.addNode(7)
	t.addNode(45)
	inOrder(t.root, func(n *node) { fmt.Println(n.val) })
}
