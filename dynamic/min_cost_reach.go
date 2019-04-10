package main

import (
	"fmt"
	"math"
)

func main() {
	m := [][]int{
		{4, 7, 8, 6, 4},
		{6, 7, 3, 9, 2},
		{3, 8, 1, 2, 4},
		{7, 1, 7, 3, 7},
		{2, 9, 8, 9, 3},
	}

	fmt.Println("mincost", minCost(m, len(m), len(m[0])))
	fmt.Println("mincost", minCostDp(m, len(m), len(m[0])))
}

func minCost(mat [][]int, m int, n int) int {
	if m == 0 || n == 0 {
		return int(math.MaxInt32)
	}

	if m == 1 && n == 1 {
		return mat[0][0]
	}

	return mat[m-1][n-1] + min(minCost(mat, m, n-1), minCost(mat, m-1, n))
}

func minCostDp(mat [][]int, m int, n int) int {
	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j > 0 {
				dp[0][j] += dp[0][j-1]
			} else if j == 0 && i > 0 {
				dp[i][0] += dp[i-1][0]
			} else if i > 0 && j > 0 {
				dp[i][j] += min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
