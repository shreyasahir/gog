package main

type node struct {
	val   int
	left  *node
	right *node
	hd    int
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

func (t *tree) remove(val int) {
	n := &node{val: val}
	remove(t.root, n)
}

func remove(curr *node, n *node) *node {
	if curr == nil {
		return nil
	}
	if curr.val > n.val {
		curr.left = remove(curr.left, n)
		return curr
	}
	if curr.val < n.val {
		curr.right = remove(curr.right, n)
		return curr
	}
	if curr.left == nil && curr.right == nil {
		curr = nil
		return curr
	}

	if curr.left == nil {
		curr = curr.right
		return curr
	}
	if curr.right == nil {
		curr = curr.left
		return curr
	}
	rightMostLeft := curr.left
	for {
		if rightMostLeft != nil && rightMostLeft.left != nil {
			rightMostLeft = rightMostLeft.left
		} else {
			break
		}
	}

	curr.val = rightMostLeft.val
	curr.right = remove(curr.right, curr)
	return curr
}
