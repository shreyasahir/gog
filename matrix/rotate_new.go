package main

import "fmt"

func main() {
	arr := [][]int{

		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotateNew(arr)
}

func rotateNew(arr [][]int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < len(arr[0]); j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}
	fmt.Println("Arr", arr)
	for _, v := range arr {
		for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
			v[i], v[j] = v[j], v[i]
		}
	}
	fmt.Println("Arr", arr)
}

func rotate(m [][]int) {

	for i := 0; i < len(m); i++ {
		for j := i + 1; j < len(m); j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}

	for _, v := range m {
		for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
			v[i], v[j] = v[j], v[i]
		}
	}
	// fmt.Printf("%v", m)
}
