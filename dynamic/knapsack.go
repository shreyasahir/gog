package main

import (
	"fmt"
	"math"
)

func main() {
	value := []int{20, 5, 10, 40, 15, 25}
	weight := []int{1, 2, 3, 8, 7, 4}
	w := 10
	fmt.Println("knapsack", knapsack(value, weight, len(value)-1, w))
}

func knapsack(value []int, weight []int, n int, w int) int {
	if w < 0 {
		return int(math.MinInt32)
	}
	if n < 0 || w == 0 {
		return 0
	}

	return max(knapsack(value, weight, n-1, w-weight[n])+value[n], knapsack(value, weight, n-1, w))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
