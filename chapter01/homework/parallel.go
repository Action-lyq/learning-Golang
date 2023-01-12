package main

import "fmt"

func main() {
	p1x, p1y := getPointXYFromInput()
	p2x, p2y := getPointXYFromInput()
	p3x, p3y := getPointXYFromInput()
	p4x, p4y := getPointXYFromInput()

	k1 := calculateK(p2y, p1y, p2x, p1x)
	k2 := calculateK(p4y, p3y, p4x, p3x)

	getResult(k1, k2)
}

func getPointXYFromInput() (float64, float64) {
	var x, y float64
	fmt.Println("录入x：")
	_, err := fmt.Scanln(&x)
	if err != nil {
		return 0, 0
	}
	fmt.Println("录入y：")
	_, err = fmt.Scanln(&y)
	if err != nil {
		return 0, 0
	}
	return x, y
}

func calculateK(y2, y1, x2, x1 float64) float64 {
	return (y2 - y1) / (x2 - x1)
}

func getResult(k1, k2 float64) {
	if k1 == k2 {
		fmt.Println("平行")
	} else {
		fmt.Println("不平行")
	}
}
