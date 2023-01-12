package main

import "fmt"

func main() {
	var name string = "刘大大"
	fmt.Println(name)
	var age int = 32
	fmt.Println(age)
	var tall float64 = 1.76
	fmt.Println(tall)

	name = "刘大大二世"
	//age = "三十二" // 类型不兼容
	//tall = age // 类型不兼容
}
