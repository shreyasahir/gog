package main

import "fmt"

func main() {
	arr := []int{3, 4, 6, 10, 11, 15}
	arr1 := []int{1, 5, 8, 12, 14, 19}
	merge(arr, arr1)
}

func merge(arr, arr1 []int) {
	m := len(arr)
	n := len(arr1)
	result := make([]int, m+n)
	for n > 0 {
		if m == 0 || arr[m-1] < arr1[n-1] {
			result[m+n-1] = arr1[n-1]
			n--
		} else {
			result[m+n-1] = arr[m-1]
			m--
		}
	}
	fmt.Println(result)
}
