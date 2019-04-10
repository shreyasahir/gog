package main

import "fmt"

func main() {
	arr := []int{7, 3, 1, 5, 4, 8}
	fmt.Println("equalsum", equalsum(arr, len(arr)))
}

func equalsum(arr []int, n int) bool {
	if n%2 != 0 {
		return false
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	return subset(arr, n, sum/2)
}

func subset(arr []int, n int, sum int) bool {
	if sum == 0 {
		return true
	}

	if n < 0 || sum < 0 {
		return false
	}

	return subset(arr, n-1, sum-arr[n]) || subset(arr, n-1, sum)
}
