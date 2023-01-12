package main

import "fmt"

func main() {
	fmt.Println("round 1")
	//for i := 0; i < 1; i++ {
	//	fmt.Println("hello, Golang")
	//}

	fmt.Println("round 2")
	//j := 0
	//for ; j < 2; j++ {
	//	fmt.Println("2 hello, Golang")
	//}

	fmt.Println("round 3")
	//k := 0
	//for ; true; k++ {
	//	fmt.Println("3 hello, Golang")
	//}

	fmt.Println("round 4")
	//l := 0
	//for l < 4 {
	//	fmt.Println("4 hello, Golang")
	//	l++
	//}

	fmt.Println("round 5")
	m := 0
	for {
		fmt.Println("5 hello, Golang")
		m++
		if m >= 5 {
			break
		}
	}

	fmt.Println("round 6")
	n := 0
	for {
		fmt.Println("6 hello, Golang", n)
		n++
		if n >= 10 {
			break
		}
		if n%2 == 0 {
			fmt.Println("被continue了")
			continue
		}
		fmt.Println("练习跳过")
	}
}
