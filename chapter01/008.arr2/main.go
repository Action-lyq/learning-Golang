package main

import "fmt"

func main() {
	lddInfo := [3]string{"刘大大", "男", "在职"}
	wsfInfo := [3]string{"王师傅", "男", "离职"}
	zleInfo := [3]string{"翟老二", "男", "在职"}
	fmt.Println(lddInfo, wsfInfo, zleInfo)

	personInfo := [3][3]string{
		{"刘大大", "男", "在职"},
		{"王师傅", "男", "离职"},
		{"翟老二", "男", "在职"},
	}
	for _, val := range personInfo {
		fmt.Println(val)
	}

	newpersonInfo := [...][3]string{
		{"刘大大", "男", "在职"},
		{"王师傅", "男", "离职"},
		{"翟老二", "男", "在职"},
	}
	for _, val := range newpersonInfo {
		fmt.Println(val)
	}

	fmt.Println("用降维的方式输出")
	for d1, d1Val := range newpersonInfo {
		for d2, d2Val := range d1Val {
			fmt.Println(d1, d1Val, d2, d2Val)
		}
	}
}
