package main

import "fmt"

func main() {
	hello()
	helloToSomeOne("刘大大")
	msg := constructHelloMessage("王师傅")
	fmt.Println(msg)
}

func hello() {
	fmt.Println("你好，Golang")
}

func helloToSomeOne(name string) {
	fmt.Println("你好", name)
}

func constructHelloMessage(name string) string {
	fmt.Println("你好", name)
	return "你好" + name
}
