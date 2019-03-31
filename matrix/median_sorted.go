package main

import "fmt"

func main() {
	m := [][]int{
		{1, 3, 5},
		{2, 6, 9},
		{3, 6, 9},
	}
	binaryMedian(m)
}

func binaryMedian(m [][]int) {
	mi := m[0][0]
	mx := 0
	r := 3
	d := 3
	for i := 0; i < r; i++ {
		if m[i][0] > mi {
			mi = m[i][0]
		}
		if m[i][d-1] > mx {
			mx = m[i][d-1]
		}
	}
	desired := (r*d + 1) / 2
	for mi < mx {
		mid := mi + (mx-mi)/2
		place := 0
		for i := 0; i < r; i++ {
			j := CountOfNum(m[i], mid)
			place = place + j
		}
		if place < desired {
			mi = mid + 1
		} else {
			mx = mid
		}
	}

	fmt.Println("mid", mi)
}

func CountOfNum(arr []int, mid int) int {
	count := 0
	for _, v := range arr {
		if v <= mid {
			count++
		}
	}
	return count
}
