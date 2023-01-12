package main

import "fmt"

func main() {
	names := []string{"刘大大", "王师傅", "翟老师"}
	fr := []float64{28, 18, 18}
	names = append(names, "小强")
	fr = append(fr, 19)

	for i, name := range names {
		if name == "王师傅" {
			fmt.Printf("%s 的体脂率是%f\n\n", name, fr[i])
		}
	}

	//var m1 map[string]int
	//m2 := map[string]int{}
	m3 := map[string]int{"刘大大": 85, "王师傅": 85, "翟老师": 80}
	//fmt.Println(m1, m2, m3)
	//m3["小强"] = 90
	fmt.Println(m3)
	m3["刘大大"] = 70
	fmt.Println(m3)
	delete(m3, "王师傅")
	fmt.Println(m3)

	for k, v := range m3 {
		fmt.Printf("名称：%s, 分数：%d\n", k, v)
	}

	xqScore, ok := m3["小强"]
	fmt.Println(xqScore, ">>>>>", ok)
	m3["小强"] = 90
	xqScore, ok = m3["小强"]
	fmt.Println(xqScore, ">>>>>", ok)
}
