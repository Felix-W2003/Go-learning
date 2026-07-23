package main

import "fmt"

func isPalindrome(x []int) bool {
	for i := 0; i < len(x)/2; i++ {
		if x[i] != x[len(x)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 1, 1}
	b := []int{1, 1, 1, 1, 1}
	c := []int{2, 1, 3, 2, 2}
	fmt.Println(isPalindrome(a), isPalindrome(b), isPalindrome(c))
}
