package main

import "fmt"

func main() {
	x := "kitten"
	y := "sitting"
	fmt.Println("edit", editDistance(x, len(x), y, len(y)))
}

func editDistance(x string, m int, y string, n int) int {
	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}
	if x[m-1] == y[n-1] {
		editDistance(x, m-1, y, n-1)
	}

	return 1 + min(editDistance(x, m-1, y, n), editDistance(x, m, y, n-1), editDistance(x, m-1, y, n-1))
}

// func editDistance(s1 string, s2 string, m int, n int) int {
// 	if m == 0 {
// 		return n
// 	}

// 	if n == 0 {
// 		return m
// 	}

// 	if s1[m-1] == s2[n-1] {
// 		return editDistance(s1, s2, m-1, n-1)
// 	} else {
// 		return 1 + min(editDistance(s1, s2, m, n-1), editDistance(s1, s2, m-1, n), editDistance(s1, s2, m-1, n-1))
// 	}
// }

// func min(a, b, c int) int {
// 	if a < b {
// 		if a < c {
// 			return a
// 		}
// 		return c

// 	}
// 	if b < c {
// 		return b
// 	}
// 	return c
// }

func min(a, b, c int) int {
	fmt.Println("a,b,c", a, b, c)
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
