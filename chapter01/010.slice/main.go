package main

import "fmt"

func main() {
	var a = "hello"
	fmt.Println(a)
	aBytes := []byte(a) // []int 不能转, []byte可以和string互转
	fmt.Println(aBytes)
	fmt.Println("修改切片内的内容")
	aBytes[0] = 'H'
	a = string(aBytes)
	fmt.Println(a)

	b := "您好"
	fmt.Println(b)
	bBytes := []rune(b)
	fmt.Println(bBytes)
	bBytes[0] = 'H'
	b = string(bBytes)
	fmt.Println(b)
}
