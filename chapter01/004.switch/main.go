package main

import "fmt"

func main() {
	var money int = 20
	var busy bool
	switch money {
	case 20:
		fmt.Println("点外卖")
		fallthrough
	case 200:
		fmt.Println("下馆子")
		if busy {
			break
		} else {
			fmt.Println("再吃点零食")
		}
	default:
		fmt.Println("容我想想")
	}
}
