package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a)
	fmt.Println("追加元素到a中，a是切片")
	//fmt.Println(append(a, 2))
	b := [0]int{}
	fmt.Println(b)
	fmt.Println("追加元素到吧，b是数组")
	//b=append(a,3)

	lddInfo := []string{"刘大大", "男", "在职"}
	fmt.Println(lddInfo)
	for i, v := range lddInfo {
		fmt.Println(i, v)
	}
	fmt.Println(lddInfo[0])

	fmt.Println("=========")
	fmt.Println("删除切片种的元素")
	delArr := []int{11, 22, 33, 44, 55}
	fmt.Println("删除前", delArr)
	fmt.Println("删除后", append(delArr[:1], delArr[1:4]...))
	fmt.Println("删除后", append(delArr[:2], delArr[3:5]...))
}
