package main

import "fmt"

func main() {
	for {
		mainFatRateBody()

		if cont := whetherContinue(); !cont {
			break
		}
	}
}

func mainFatRateBody() {
	bmi, age, sex, sexWeight := getMaterialsFromInput()

	fatRate := calcFatRate(bmi, age, sexWeight)

	getHealthinessSuggestions(sex, age, fatRate)
}

func getMaterialsFromInput() (float64, int, string, int) {
	// 录入各项
	var weight = 97.5
	fmt.Print("体重（千克）：")
	_, err := fmt.Scanln(&weight)
	if err != nil {
		return 0, 0, "", 0
	}

	var tall = 1.76
	fmt.Print("身高（米）：")
	_, err = fmt.Scanln(&tall)
	if err != nil {
		return 0, 0, "", 0
	}

	var bmi = weight / (tall * tall)

	var age = 32
	fmt.Print("年龄：")
	_, err = fmt.Scanln(&age)
	if err != nil {
		return 0, 0, "", 0
	}

	var sex = "男"
	fmt.Print("性别（男/女）：")
	_, err = fmt.Scanln(&sex)
	if err != nil {
		return 0, 0, "", 0
	}

	var sexWeight int
	if sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	return bmi, age, sex, sexWeight
}

func calcFatRate(bmi float64, age int, sexWeight int) float64 {
	// 计算体脂率
	var fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	fmt.Println("体脂率为：", fatRate)
	return fatRate
}

func getHealthinessSuggestions(sex string, age int, fatRate float64) {
	// 得到健康提示
	if sex == "男" {
		// todo 编写男性的体脂率与体脂状态表
		if age >= 18 && age <= 39 {
			if fatRate <= 0.1 {
				fmt.Println("偏瘦")
			} else if fatRate > 0.1 && fatRate <= 0.16 {
				fmt.Println("标准")
			} else if fatRate > 0.16 && fatRate <= 0.21 {
				fmt.Println("偏胖")
			} else if fatRate > 0.21 && fatRate <= 0.26 {
				fmt.Println("肥胖")
			} else {
				fmt.Println("严重肥胖")
			}
		} else if age >= 40 && age <= 59 {

		} else if age >= 60 {

		} else {
			fmt.Println("不计算未成年人的体脂率")
		}
	} else {
		// todo 编写女性的体脂率与体脂状态表
	}
}

func whetherContinue() bool {
	var whetherContinue string
	fmt.Print("是否录入下一个（y/n）?")
	_, err := fmt.Scanln(&whetherContinue)
	if err != nil {
		return false
	}
	if whetherContinue != "y" {
		return false
	}
	return true
}
