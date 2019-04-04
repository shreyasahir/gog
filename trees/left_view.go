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
	printLeft(t.root)
}

func printLeft(n *node) {
	if n == nil {
		return
	}
	var queue []*node
	queue = append(queue, n)

	for len(queue) > 0 {

		level := 0
		for level < len(queue) {
			curr := queue[0]
			queue = queue[1:]
			if level == 0 {
				fmt.Println(curr.val)
			}
			if curr.left != nil {
				queue = append(queue, curr.left)
			}

			if curr.right != nil {
				queue = append(queue, curr.right)
			}
			level++
		}
	}

	//fmt.Printf("%v", hash)

}
