package main

import "fmt"

func main() {
	t := &tree{}
	t.addNode(8)
	t.addNode(6)
	t.addNode(5)
	t.addNode(15)
	t.addNode(29)
	t.addNode(27)
	t.addNode(2)
	t.addNode(44)
	t.addNode(7)
	t.addNode(45)
	reverLevelOrder(t.root)
}

func reverLevelOrder(n *node) {
	if n == nil {
		return
	}
	var queue, stack []*node

	queue = append(queue, n)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		stack = append(stack, curr)
		if curr.right != nil {
			queue = append(queue, curr.right)
		}
		if curr.left != nil {
			queue = append(queue, curr.left)
		}
	}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println("curr", curr.val)
	}
}
