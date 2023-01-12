package main

import (
	"fmt"
	"math"
)

func main() {
	var hello string = "hello"
	var world = "world"
	fmt.Println(hello, world)
	xs := 3.1415926
	fmt.Println(xs)

	var int3, int4 uint = 3, 4
	fmt.Println(int3, int4)

	var (
		int1, int2 = 1, 2
	)
	fmt.Println(int1, int2)

	// 类型推断
	var ho, ver float64 = 3.14, 4
	var sc = ho * ver
	fmt.Println(ho * ver)
	fmt.Println(sc)

	//设置默认值
	var newname string
	fmt.Println(newname)

	//强制转换
	var int5 uint = math.MaxUint64
	fmt.Println(int5)
	var int6 = int(int5)
	fmt.Println(int6)
}
