package main

import "fmt"

func main() {
	m := 40000
	n := 40000
	ops := [][]int{}
	fmt.Println(maxCount(m, n, ops))
}

func maxCount(m int, n int, ops [][]int) int {

	result := make([][]int, m)
	for k := range result {
		result[k] = make([]int, n)
	}

	for _, v := range ops {
		if v != nil {
			for i := 0; i < v[0]; i++ {
				for j := 0; j < v[1]; j++ {
					result[i][j]++
				}
			}
		}
	}
	//fmt.Println("res",result)
	//x := maxOfMatrix(result)
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if result[i][j] == result[0][0] {
				count++
			}
		}
	}
	return count
}

func maxOfMatrix(matrix [][]int) int {
	maximum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			maximum = max(maximum, matrix[i][j])
		}
	}
	return maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
