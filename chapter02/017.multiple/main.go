package main

import "fmt"

func main() {
	//bmis := []float64{1.2, 3.3, 4.5}
	avgBmi := calculateAvg(1, 2, 3, 4, 45, 6, 6, 7)
	fmt.Println(avgBmi)
	getScoresOfStudent()
}

func calculateAvg(bmis ...float64) (avgBmi float64) {
	var total float64 = 0
	for _, item := range bmis {
		total += item
	}
	return total / float64(len(bmis))
}

func getScoresOfStudent() (chinese int, math int, english int, physics int, nature int) {
	return 0, 0, 0, 0, 0
}
