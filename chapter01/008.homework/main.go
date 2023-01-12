package main

import "fmt"

func main() {
	// 创建一个一维数组，反转它
	var arr = [3]int{1, 3, 5}
	var arrLen = len(arr)
	var reversalArr = [3]int{}
	for _, val := range arr {
		arrLen--
		reversalArr[arrLen] = val
	}
	fmt.Println(reversalArr)

	// todo 用多维数组写一个日历表，需要考虑每个月的天数不同，需要考虑是平年还是闰年来专门处理二月

	// todo【提升篇】日历按照一周的宽度显示（第一列为周一），且每个日期匹配到对应的列

}
