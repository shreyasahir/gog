package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 4
	fmt.Println("count", countAndSay(n))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	curr := "1"
	for c := 2; c <= n; c++ {
		curr = getNext(curr)
	}
	return curr
}

func getNext(s string) string {
	ret := ""
	i := 0
	for i < len(s) {
		j := i + 1
		for j < len(s) {
			if s[j-1] != s[j] {
				break
			} else {
				j++
			}
		}
		ret = ret + strconv.Itoa(j-i) + string(s[i])
		i = j
	}
	return ret
}
