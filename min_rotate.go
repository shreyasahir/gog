package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("min", rotateMin(arr))
}

func rotateMin(arr []int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2
		if arr[mid-1] < arr[mid] && arr[mid] > arr[mid+1] {
			return arr[mid]
		} else if arr[mid] < arr[mid+1] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
