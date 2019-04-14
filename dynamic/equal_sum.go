package main

import "fmt"

func main() {
	arr := []int{7, 3, 1, 5, 4, 8}
	fmt.Println("equalsum", equalsum(arr, len(arr)))
}

func equalsum(arr []int, n int) bool {

	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	// return (sum&1) == 0 && subset(arr, n-1, sum/2)
	return (sum&1) == 0 && subsetDp(arr, n-1, sum/2)
}

func subset(arr []int, n int, sum int) bool {
	fmt.Println("sum", sum)
	if sum == 0 {
		return true
	}

	if n < 0 || sum < 0 {
		return false
	}

	return subset(arr, n-1, sum-arr[n]) || subset(arr, n-1, sum)
}

func subsetDp(arr []int, n int, sum int) bool {
	dp := make([][]bool, n+1)
	for k := range dp {
		dp[k] = make([]bool, sum+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if arr[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-arr[i-1]]
			}
		}
	}
	return dp[n][sum]
}
