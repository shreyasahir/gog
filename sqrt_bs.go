package main

import "fmt"

func main() {
	n := 11
	fmt.Println(sqrt(n))
}

func sqrt(n int) int {
	low := 0
	high := n

	for low <= high {
		mid := low + (high-low)/2
		if mid*mid == n {
			return mid
		} else if mid*mid > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		m := l + (r-l)/2
		if m*m == x {
			return m
		} else if m*m > x {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return r
}
