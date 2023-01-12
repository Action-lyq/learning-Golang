package main

import "fmt"

func main() {
	var circle1, circle2, pi float64 = 1, 2.1, 3.141
	circle1Area := circle1 * pi
	circle2Area := circle2 * pi
	fmt.Printf("%.3f", circle1Area+circle2Area)
}
