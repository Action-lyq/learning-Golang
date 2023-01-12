package main

import "fmt"

func main() {
	for {
		var name [3]string
		var weight [3]float64
		var tall [3]float64
		var bmi [3]float64
		var age [3]int
		var sex [3]string
		var sexWeight [3]int
		var fatRate [3]float64

		for i := 0; i < 3; i++ {
			fmt.Print("姓名：")
			_, err := fmt.Scanln(&name[i])
			if err != nil {
				return
			}

			fmt.Print("体重（千克）：")
			_, err = fmt.Scanln(&weight[i])
			if err != nil {
				return
			}

			fmt.Print("身高（米）：")
			_, err = fmt.Scanln(&tall[i])
			if err != nil {
				return
			}

			bmi[i] = weight[i] / (tall[i] * tall[i])

			fmt.Print("年龄：")
			_, err = fmt.Scanln(&age[i])
			if err != nil {
				return
			}

			fmt.Print("性别（男/女）：")
			_, err = fmt.Scanln(&sex[i])
			if err != nil {
				return
			}

			if sex[i] == "男" {
				sexWeight[i] = 1
			} else {
				sexWeight[i] = 0
			}
			fatRate[i] = (1.2*bmi[i] + 0.23*float64(age[i]) - 5.4 - 10.8*float64(sexWeight[i])) / 100
			fmt.Print("姓名：", name[i], "BMI：", bmi[i], "体脂率：", fatRate[i], "建议：")
			if sex[i] == "男" {
				// todo 编写男性的体脂率与体脂状态表
				if age[i] >= 18 && age[i] <= 39 {
					if fatRate[i] <= 0.1 {
						fmt.Println("偏瘦")
					} else if fatRate[i] > 0.1 && fatRate[i] <= 0.16 {
						fmt.Println("标准")
					} else if fatRate[i] > 0.16 && fatRate[i] <= 0.21 {
						fmt.Println("偏胖")
					} else if fatRate[i] > 0.21 && fatRate[i] <= 0.26 {
						fmt.Println("肥胖")
					} else {
						fmt.Println("严重肥胖")
					}
				} else if age[i] >= 40 && age[i] <= 59 {

				} else if age[i] >= 60 {

				} else {
					fmt.Println("不计算未成年人的体脂率")
				}
			} else {
				// todo 编写女性的体脂率与体脂状态表
			}
		}

		fmt.Println("总人数：", len(name), "平均体脂率：", (fatRate[0]+fatRate[1]+fatRate[2])/3)

		var whetherContinue string
		fmt.Print("是否录入下一个（y/n）?")
		_, err := fmt.Scanln(&whetherContinue)
		if err != nil {
			return
		}
		if whetherContinue != "y" {
			break
		}
	}
}
