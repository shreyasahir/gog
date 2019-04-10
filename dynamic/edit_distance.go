package main

import "fmt"

func main() {
	x := "kitten"
	y := "sitting"
	fmt.Println("edit", editDistance(x, y, len(x), len(y)))
	fmt.Println("edit", editDistanceDp(x, y))

}

func editDistance(s1 string, s2 string, m int, n int) int {
	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	if s1[m-1] == s2[n-1] {
		return editDistance(s1, s2, m-1, n-1)
	} else {
		return 1 + min(editDistance(s1, s2, m, n-1), editDistance(s1, s2, m-1, n), editDistance(s1, s2, m-1, n-1))
	}
}

func editDistanceDp(x string, y string) int {
	m := len(x)
	n := len(y)

	dp := make([][]int, m+1)
	for k := range dp {
		dp[k] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if x[i-1] == y[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}

		}
	}
	return dp[m][n]
}
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
