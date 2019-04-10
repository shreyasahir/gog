package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}
	fmt.Println("longIncrSub", longIncrSub(arr, 0, len(arr)-1, int(math.MinInt32)))
	fmt.Println("longIncrSubIter", longIncrSubIter(arr, len(arr)))

	nums := []int{8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11}
	fmt.Println("longIncrSubMaxSum", longIncrSubMaxSum(nums, len(nums)))
}

func longIncrSub(arr []int, i int, n int, prev int) int {
	if n == i {
		return 0
	}
	excld := longIncrSub(arr, i+1, n, prev)

	incld := 0
	if arr[i] > prev {

		incld = longIncrSub(arr, i+1, n, arr[i]) + 1
	}

	return max(excld, incld)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longIncrSubIter(arr []int, n int) int {
	l := make([]int, n)
	l[0] = 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && l[j] > l[i] {
				l[i] = l[j]
			}
		}
		l[i]++
	}
	fmt.Println("arr", l)
	return 1
}

func longIncrSubMaxSum(arr []int, n int) int {
	l := make([]int, n)
	l[0] = arr[0]
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && l[j] > l[i] {
				l[i] = l[j]
			}
		}
		l[i] += arr[i]
	}
	fmt.Println("arr", l)
	return 1
}
