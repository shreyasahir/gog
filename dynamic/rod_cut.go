package main

import "fmt"

func main() {
	length := []int{1, 2, 3, 4, 5, 6, 7, 8}
	price := []int{1, 5, 8, 9, 10, 17, 17, 20}
	rod := 4

	fmt.Println("rodcut", rodCutDp(length, price, rod))
}

func rodCut(length []int, price []int, n int) int {
	if n == 0 {
		return 0
	}
	maxCost := 0
	cost := 0
	for i := 1; i <= n; i++ {
		cost = price[i-1] + rodCut(length, price, n-i)
		if cost > maxCost {
			maxCost = cost
		}
	}
	return maxCost
}

func rodCutDp(length []int, price []int, n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = max(dp[i], price[j-1]+dp[i-j])
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
