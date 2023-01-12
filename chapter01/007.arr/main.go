package main

import "fmt"

func main() {
	var age int = 32
	fmt.Println(age)
	var ages [5]int = [5]int{29, 30, 31, 32, 33}
	var ages2 = [5]int{1, 3, 5, 7, 9}
	ages3 := [5]int{2, 4, 6, 8, 10}
	fmt.Println(ages, ages2, ages3)
	ages2 = ages3

	ages4 := [...]int{2, 4, 6, 8} //由go进行计算长度
	fmt.Println(ages4)

	var ages5 [3]int
	ages5[0] = 100
	fmt.Println(ages5)

	for i := 0; i < len(ages5); i++ {
		fmt.Println(ages5[i])
	}

	for j := range ages5 {
		fmt.Println(ages5[j])
	}
	for key, val := range ages5 {
		fmt.Println(ages5[key], "===>", key, "->", val)
	}
}
