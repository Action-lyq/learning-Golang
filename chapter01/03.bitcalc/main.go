package main

import "fmt"

func main() {
	var a, b int = 100, 31
	fmt.Println(a ^ b)
	fmt.Println(b ^ a)

	var arr = []int{4, 3, 4, 5, 13, 7, 3, 5, 13}
	var result int = -1
	for _, item := range arr {
		if result < 0 {
			result = item
		} else {
			result = result ^ item
		}
	}
	fmt.Println(result)
	fmt.Println(5 ^ 5)
	fmt.Println(4 ^ 4 ^ 3)
}
