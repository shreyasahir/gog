package main

import "fmt"

func main() {
	matrix := [][]int{
		{0, 0, 0, 1},
		{0, 1, 1, 1},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
	}

	fmt.Println(rowWithMax1(matrix))
}

func rowWithMax1(matrix [][]int) int {
	r := len(matrix)
	c := len(matrix[0])
	maxRow := 0
	maximum := -1

	for i := 0; i < r; i++ {
		index := first(matrix[i], 0, c-1)
		if index != -1 && c-index > maximum {
			maximum = c - index
			maxRow = i
		}
	}
	return maxRow
}

func first(arr []int, start int, end int) int {
	if end >= start {
		mid := start + (end-start)/2
		if (mid == 0 || arr[mid-1] == 0) && arr[mid] == 1 {
			return mid
		} else if arr[mid] == 0 {
			return first(arr, mid+1, end)
		} else {
			return first(arr, start, mid-1)
		}
	}
	return -1
}
