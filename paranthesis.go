package main

import "fmt"

func main() {
	n := 3
	s := ""
	paranthesis(n, n, s)
}

func paranthesis(left int, right int, s string) {
	if left == 0 && right == 0 {
		fmt.Println(s)
	}
	if left > right {
		return
	}
	if left > 0 {
		paranthesis(left-1, right, s+"(")
	}

	if right > 0 {
		paranthesis(left, right-1, s+")")
	}
}
