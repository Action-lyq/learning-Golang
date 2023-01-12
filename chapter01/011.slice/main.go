package main

import "fmt"

func main() {
	a := make([]int, 0, 2)
	fmt.Println("len:", len(a), "cap:", cap(a))
	a = append(a, 1)
	a = append(a, 1)
	a = append(a, 1)
	fmt.Println("len:", len(a), "val:", a, "cap:", cap(a))
}
