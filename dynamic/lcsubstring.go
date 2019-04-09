package main

import "fmt"

func main() {
	x := "ABC"
	y := "BABA"
	fmt.Println("lcsubstring", lcsubstring(x, len(x), y, len(y)))
}

func lcsubstring(x string, m int, y string, n int) string {
	dp := make([][]int, m+1)
	maxLength := 0
	endhere := m
	for k := range dp {
		dp[k] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = 0
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i-1] == y[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if maxLength < dp[i][j] {
					maxLength = dp[i][j]
					endhere = i
				}
			}
		}
	}
	return x[:endhere]
}
