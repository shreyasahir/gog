package main

import "fmt"

func main() {
	A := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	B := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	multiply(A, len(A), len(A[0]), B, len(B), len(B[0]))
}

func multiply(a [][]int, r1 int, c1 int, b [][]int, r2 int, c2 int) {
	if c1 != r2 {
		return
	}
	result := make([][]int, r1)
	for k := range result {
		result[k] = make([]int, c2)
	}

	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			for k := 0; k < r2; k++ {
				result[i][j] = result[i][j] + a[i][k]*b[k][j]
			}
		}
	}
	fmt.Println("multi", result)
}
