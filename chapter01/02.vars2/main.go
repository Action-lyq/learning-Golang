package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "ldd"
	{
		val, err := strconv.Atoi(name)
		fmt.Println(val, err)
	}
	age := 30
	fmt.Printf("%p\n", &age)
	age, time := 32, "时间"
	fmt.Printf("%p\n", &age)
	//age = int(name)
	fmt.Println(name, age, time)
}
