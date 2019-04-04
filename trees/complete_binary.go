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
	fmt.Println("complete", isCompleteBinary(t.root))
}

func isCompleteBinary(root *node) bool {
	if root == nil {
		return false
	}
	var queue []*node
	flag := false
	queue = append(queue, root)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if flag == true && (curr.left != nil || curr.right != nil) {
			return false
		}

		if curr.left == nil && curr.right != nil {
			return false
		}

		if curr.left != nil {
			queue = append(queue, curr.left)
		} else {
			flag = true
		}

		if curr.right != nil {
			queue = append(queue, curr.right)
		} else {
			flag = true
		}

	}
	return true
}
