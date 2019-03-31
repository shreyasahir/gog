package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println("matrix before", matrix)
	rotateMatrix(matrix)
	fmt.Println("matrix after", matrix)
}

func rotateMatrix(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = tmp
		}
	}
}
