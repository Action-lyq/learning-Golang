package main

import "fmt"

func main() {
	var fruit string = "6 apples"
	var watermelon bool
	if watermelon {
		fruit = "1 apples"
	} else {
		fruit = "7 apples"
	}
	fmt.Println("buy: ", fruit)
}
