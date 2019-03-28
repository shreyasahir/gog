package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{10, 2, -2, -20, 10}
	fmt.Println(kadane(arr))
}

func kadane(a []int) int {
	n := len(a)
	maxSoFar := int(math.MinInt32)
	MaxEndingHere := a[0]

	for i := 1; i < n; i++ {
		// MaxEndingHere = MaxEndingHere + a[i]
		MaxEndingHere = max(a[i], MaxEndingHere+a[i])
		maxSoFar = max(MaxEndingHere, maxSoFar)
	}
	return maxSoFar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
