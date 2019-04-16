package main

import "fmt"

func main() {
	s := []int{1, 1, 2}
	permute(s, 0, len(s)-1)
}

func permutation(a []int, low int, result *[][]int) {
	if low == len(a)-1 {
		t := make([]int, len(a))
		copy(t, a)
		*result = append(*result, t)
		return
	} else {
		for i := low; i < len(a); i++ {
			a[i], a[low] = a[low], a[i]
			permutation(a, low+1, result)
			a[i], a[low] = a[low], a[i]
		}
	}
}

func permute(a []int, l, r int) {

	if l == r {
		fmt.Println(a)
	} else {
		for i := l; i < r+1; i++ {
			a[i], a[l] = a[l], a[i]
			permute(a, l+1, r)
			a[i], a[l] = a[l], a[i]
		}
	}

}
