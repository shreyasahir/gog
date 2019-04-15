package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{1, 2, 3, 4}
	fmt.Println("coinChange", coinChange(s, len(s)-1, 15))
}

func coinChange(s []int, n int, sum int) int {
	if sum == 0 {
		return 0
	}
	if sum < 0 {
		return int(math.MaxInt32)
	}
	coins := int(math.MaxInt32)
	for i := 0; i <= n; i++ {
		res := coinChange(s, n, sum-s[i])
		if res != int(math.MaxInt32) {
			coins = min(coins, res+1)
		}
	}
	return coins
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
