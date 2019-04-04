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
	printSpecific(t.root)
}

func printSpecific(root *node) {
	if root == nil {
		return
	}
	var queue1, queue2 []*node
	fmt.Println(root.val)
	queue1 = append(queue1, root.left)
	queue2 = append(queue2, root.right)
	for len(queue1) > 0 {
		n := len(queue1)
		for n > 0 {
			n--
			curr := queue1[0]
			queue1 = queue1[1:]
			fmt.Println(curr.val)
			if curr.left != nil {
				queue1 = append(queue1, curr.left)
			}
			if curr.right != nil {
				queue1 = append(queue1, curr.right)
			}

			curr = queue2[0]
			queue2 = queue2[1:]
			fmt.Println(curr.val)
			if curr.right != nil {
				queue2 = append(queue2, curr.right)
			}
			if curr.left != nil {
				queue2 = append(queue2, curr.left)
			}

		}
	}
}
