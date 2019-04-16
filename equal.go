package main

import "fmt"

func main() {
	s := []int{3, 4, 7, 1, 2, 9, 8}
	equal(s)
}

func equal(arr []int) {
	hash := make(map[int][]int)
	var sum int
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			sum = arr[i] + arr[j]
			if _, ok := hash[sum]; ok {
				storedValue := hash[sum]
				fmt.Println(storedValue[0], storedValue[1], arr[i], arr[j])
			} else {
				sum = arr[i] + arr[j]
				hash[sum] = []int{arr[i], arr[j]}
			}
		}
	}
}
