package main

import "fmt"

func main() {
	x := "ABCBDAB"
	y := "BDCABA"
	fmt.Println(lcs(x, len(x), y, len(y)))
	fmt.Println("dynamic", lcsdynamic(x, y))
	fmt.Println(lcsAll(x, len(x), y, len(y)))
}

func lcs(x string, m int, y string, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if x[m-1] == y[n-1] {
		return lcs(x[:m-1], m-1, y[:n-1], n-1) + 1
	}

	return max(lcs(x, m, y[:n-1], n-1), lcs(x[:m-1], m-1, y, n))

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lcsdynamic(x, y string) int {
	m := len(x)
	n := len(y)
	dp := make([][]int, m+1)
	for k := range dp {
		dp[k] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i-1] == y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func lcsAll(x string, m int, y string, n int) string {
	if m == 0 || n == 0 {
		return ""
	}
	if x[m-1] == y[n-1] {
		return lcsAll(x[:m-1], m-1, y[:n-1], n-1) + string(x[m-1])
	}

	return max(lcsAll(x, m, y[:n-1], n-1), lcsAll(x[:m-1], m-1, y, n))

}
