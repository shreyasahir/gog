package main

import "fmt"

func main() {
	arr := []int{3, 2, 0, 1}
	// Arr[i]
	// Arr[Arr[i]]
	//[0,1]
	arraychange(arr)
}

func arraychange(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		arr[i] += (arr[arr[i]] % n) * n
	}
	fmt.Println("arr", arr)
	for i := 0; i < n; i++ {
		arr[i] = arr[i] / n
	}
	fmt.Println("arr", arr)

}
