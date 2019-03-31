package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println("matrix", matrix)
	rotateMatrix(matrix)
}

func rotateMatrix(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for top < bottom && left < right {
		prev := matrix[top+1][left]

		for i := left; i < right+1; i++ {
			tmp := matrix[top][i]
			matrix[top][i] = prev
			prev = tmp
		}
		top++

		for i := top; i < bottom+1; i++ {
			tmp := matrix[i][right]
			matrix[i][right] = prev
			prev = tmp
		}
		right--

		for i := right; i > left-1; i-- {
			tmp := matrix[bottom][i]
			matrix[bottom][i] = prev
			prev = tmp
		}
		bottom--

		for i := bottom; i > top-1; i-- {
			tmp := matrix[i][left]
			matrix[i][left] = prev
			prev = tmp
		}
		left++
	}
	fmt.Println("mat", matrix)
}
