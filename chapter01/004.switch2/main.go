package main

import (
	"fmt"
	"reflect"
)

func main() {
	var money interface{} = 10
	switch newMoney := money.(type) {
	case int:
		fmt.Println("是int", newMoney+3.0)
	case int64:
		//fmt.Println("是int64", newMoney+3.012)
	case float64:
		fmt.Println("是float64", newMoney)
	case float32:
		fmt.Println("是float32")
	default:
		fmt.Println("是未知类型")
	}
	fmt.Println("结束", money, reflect.TypeOf(money))
	money = 3
	fmt.Println("结束", money, reflect.TypeOf(money))
}
