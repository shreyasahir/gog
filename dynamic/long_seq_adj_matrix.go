package main

import "fmt"

func main() {
	m := [][]int{
		{10, 13, 14, 21, 23},
		{11, 9, 22, 2, 3},
		{12, 8, 1, 5, 4},
		{15, 24, 7, 6, 20},
		{16, 17, 18, 19, 25},
	}

	maxStr := ""
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			str := findLongSeq(m, i, j)
			fmt.Println("str", str)
			if len(str) > len(maxStr) {
				maxStr = str
			}
		}
	}
	fmt.Println("maxstr", maxStr)
}

func findLongSeq(m [][]int, i int, j int) string {
	if i < 0 || j < 0 || i >= len(m) || j >= len(m[0]) {
		return ""
	}
	var path string
	if i > 0 && m[i-1][j]-m[i][j] == 1 {
		path = findLongSeq(m, i-1, j)
	}

	if j+1 < len(m) && m[i][j+1]-m[i][j] == 1 {
		path = findLongSeq(m, i, j+1)
	}

	if i+1 < len(m) && m[i+1][j]-m[i][j] == 1 {
		path = findLongSeq(m, i+1, j)
	}

	if j > 0 && m[i][j-1]-m[i][j] == 1 {
		path = findLongSeq(m, i, j-1)
	}

	fmt.Println("path", path)
	if path != "" {
		return string(m[i][j]) + path
	}
	return string(m[i][j])
}

// // recurse top cell if its value is +1 of value at (i, j)
// if (i > 0 && M[i - 1][j] - M[i][j] == 1) {
// 	path = findLongestPath(M, i - 1, j);
// }

// // recurse right cell if its value is +1 of value at (i, j)
// if (j + 1 < M.length && M[i][j + 1] - M[i][j] == 1) {
// 	path = findLongestPath(M, i, j + 1);
// }

// // recurse bottom cell if its value is +1 of value at (i, j)
// if (i + 1 < M.length && M[i + 1][j] - M[i][j] == 1) {
// 	path = findLongestPath(M, i + 1, j);
// }

// // recurse left cell if its value is +1 of value at (i, j)
// if (j > 0 && M[i][j - 1] - M[i][j] == 1) {
// 	path = findLongestPath(M, i, j - 1);
// }
