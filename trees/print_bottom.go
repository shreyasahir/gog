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
	printBottom(t.root)
}

type vector struct {
	level int
	key   int
}

func printBottom(root *node) {
	hash := make(map[int][]int)
	if root == nil {
		return
	}
	var queue []*node
	root.hd = 0
	queue = append(queue, root)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		hash[curr.hd] = append(hash[curr.hd], curr.val)

		if curr.left != nil {
			curr.left.hd = curr.hd - 1
			queue = append(queue, curr.left)
		}

		if curr.right != nil {
			curr.right.hd = curr.hd + 1
			queue = append(queue, curr.right)
		}
	}

	fmt.Println("bottom", hash)
}
