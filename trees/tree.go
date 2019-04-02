package main

type node struct {
	val   int
	left  *node
	right *node
}

type tree struct {
	root *node
}

func (t *tree) addNode(value int) {
	n := &node{val: value}
	if t.root == nil {
		t.root = n
	} else {
		addNode(t.root, n)

	}
}

func addNode(n1 *node, n2 *node) error {
	switch {
	case n1.val < n2.val:
		{
			if n1.right == nil {
				n1.right = n2
				return nil
			}
			return addNode(n1.right, n2)
		}
	case n1.val > n2.val:
		{
			if n1.left == nil {
				n1.left = n2
				return nil
			}
			return addNode(n1.left, n2)

		}
	}
	return nil
}

func inOrder(n *node, f func(*node)) {
	if n != nil {
		inOrder(n.left, f)
		f(n)
		inOrder(n.right, f)
	}
}

func preOrder(n *node, f func(*node)) {
	if n != nil {
		f(n)
		preOrder(n.left, f)
		preOrder(n.right, f)
	}
}

func postOrder(n *node, f func(*node)) {
	if n != nil {
		postOrder(n.left, f)
		postOrder(n.right, f)
		f(n)
	}
}
