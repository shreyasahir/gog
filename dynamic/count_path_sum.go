package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := [][]int{
		{4, 7, 1, 6},
		{5, 7, 3, 9},
		{3, 2, 1, 2},
		{7, 1, 6, 3},
	}
	hash := make(map[string]int)
	fmt.Println("countPath", countPath(m, len(m)-1, len(m[0])-1, 25))
	fmt.Println("countPathDp", countPathDp(m, len(m)-1, len(m[0])-1, 25, &hash))

}

func countPath(mat [][]int, m int, n int, sum int) int {
	if sum < 0 {
		return 0
	}
	if m == 0 && n == 0 {
		if mat[0][0] == sum {
			return 1
		}
		return 0
	}
	if m == 0 {
		return countPath(mat, m, n-1, sum-mat[m][n])
	}

	if n == 0 {
		return countPath(mat, m-1, n, sum-mat[m][n])
	}
	return countPath(mat, m, n-1, sum-mat[m][n]) + countPath(mat, m-1, n, sum-mat[m][n])
}

func countPathDp(mat [][]int, m int, n int, sum int, hash *map[string]int) int {
	if sum < 0 {
		return 0
	}
	if m == 0 && n == 0 {
		if mat[0][0] == sum {
			return 1
		}
		return 0
	}
	key := strconv.Itoa(m) + "|" + strconv.Itoa(n)
	_, ok := (*hash)[key]
	if !ok {
		if m == 0 {
			(*hash)[key] = countPath(mat, m, n-1, sum-mat[m][n])
		}

		if n == 0 {
			(*hash)[key] = countPath(mat, m-1, n, sum-mat[m][n])
		}
		(*hash)[key] = countPath(mat, m, n-1, sum-mat[m][n]) + countPath(mat, m-1, n, sum-mat[m][n])
	}

	return (*hash)[key]
}
