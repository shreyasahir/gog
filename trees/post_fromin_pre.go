package main

import "fmt"

func main() {
	inorder := []int{4, 2, 5, 1, 3, 6}
	preorder := []int{1, 2, 4, 5, 3, 6}
	root := printPostOrder(inorder, preorder)
	postOrder(root, func(n *node) { fmt.Println(n.val) })
}

func printPostOrder(inorder []int, preorder []int) *node {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := preorder[0]
	preorder = preorder[1:]
	index := findIndex(inorder, root)
	if index == -1 {
		return nil
	}
	inorderLeft := inorder[:index]
	inorderRight := inorder[index+1:]
	preorderLeft := preorder[:len(inorderLeft)]
	preorderRight := preorder[len(inorderLeft):]
	t := &tree{}
	t.root = &node{root, printPostOrder(inorderLeft, preorderLeft), printPostOrder(inorderRight, preorderRight)}
	return t.root
}

func findIndex(arr []int, root int) int {
	for k, v := range arr {
		if v == root {
			return k
		}
	}
	return -1
}
