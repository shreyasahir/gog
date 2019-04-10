package main

import "fmt"

func main() {
	arr := []int{4, 2, 5, 9, 7, 6, 10, 3, 1}
	lbs(arr)
}

// int n = A.length - 1;

// 		I[0] = 1;
// 		for (int i = 1; i <= n; i++) {
// 			for (int j = 0; j < i; j++) {
// 				if (A[j] < A[i] && I[j] > I[i])
// 					I[i] = I[j];
// 			}
// 			I[i]++;
// 		}

// 		D[n] = 1;
// 		for (int i = n - 1; i >= 0; i--) {
// 			for (int j = n; j > i; j--) {
// 				if (A[j] < A[i] && D[j] > D[i])
// 					D[i] = D[j];
// 			}
// 			D[i]++;
// 		}

func lbs(arr []int) {
	l := make([]int, len(arr))
	d := make([]int, len(arr))
	l[0] = 1

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && l[j] > l[i] {
				l[i] = l[j]
			}
		}
		l[i]++
	}

	d[len(arr)-1] = 1
	for i := len(arr) - 2; i >= 0; i-- {
		for j := len(arr) - 1; j > i; j-- {
			if arr[i] > arr[j] && d[j] > d[i] {
				d[i] = d[j]
			}
		}
		d[i]++
	}
	fmt.Println("l", l)
	fmt.Println("d", d)
	lbs := 1
	for i := 0; i <= len(arr)-1; i++ {
		lbs = max(lbs, l[i]+d[i]-1)
	}
	fmt.Println("lbs", lbs)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
