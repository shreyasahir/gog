package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 9, 10, 1, 30, 40}
	fmt.Println("maxValue", maxValue(arr, len(arr)))
}

func maxValue(arr []int, n int) int {
	minimum := int(math.MinInt32)
	l1 := make([]int, n+1)
	l2 := make([]int, n)
	l3 := make([]int, n-1)
	l4 := make([]int, n-2)
	for i := 0; i < n-2; i++ {
		l1[i] = minimum
		l3[i] = minimum
		l3[i] = minimum
		l4[i] = minimum
	}
	l2[n-1] = minimum
	l1[n-2] = minimum
	l1[n] = minimum

	for i := n - 1; i >= 0; i-- {
		l1[i] = max(l1[i+1], arr[i])
	}

	for i := n - 2; i >= 0; i-- {
		l2[i] = max(l2[i+1], l1[i+1]-arr[i])
	}

	for i := n - 3; i >= 0; i-- {
		l3[i] = max(l3[i+1], l2[i+1]+arr[i])
	}

	for i := n - 4; i >= 0; i-- {
		l4[i] = max(l4[i+1], l3[i+1]-arr[i])
	}
	return l4[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
