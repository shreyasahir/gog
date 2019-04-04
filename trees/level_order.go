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
	levelOrder(t.root)
}

func levelOrder(node *node) {
	if node == nil {
		return
	}
	level := 1

	for printLevel(node, level) {
		fmt.Println("level", level)
		level++
	}
}

func printLevel(node *node, level int) bool {
	if node == nil {
		return false
	}
	if level == 1 {
		fmt.Println("node", node.val)
		return true
	}
	left := printLevel(node.left, level-1)
	right := printLevel(node.right, level-1)
	return left || right
}
