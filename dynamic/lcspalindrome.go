package main

import "fmt"

func main() {
	x := "ABBDCACB"
	fmt.Println("lcspalindrome", lcspalindrome(x, 0, len(x)-1))
}

func lcspalindrome(x string, m int, n int) int {
	if m > n {
		return 0
	}

	if m == n {
		return 1
	}

	if x[m] == x[n] {
		return lcspalindrome(x, m+1, n-1) + 2
	}
	return max(lcspalindrome(x, m, n-1), lcspalindrome(x, m+1, n))

}

func lcsPalindromedp(x string, m int, n int, hash *map[string]int) int {

	if m > n {
		return 0
	}
	if m == n {
		return 1
	}

	key := string(m) + "|" + string(n)
	if _, ok := (*hash)[key]; !ok {

		if x[m] == x[n] {
			(*hash)[key] = lcsPalindromedp(x, m+1, n-1, hash) + 2
		}

		(*hash)[key] = max(lcsPalindromedp(x, m, n-1, hash), lcsPalindromedp(x, m+1, n, hash))
	}
	return (*hash)[key]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
