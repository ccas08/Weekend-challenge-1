package bmiCalculator

import (
	"fmt"
)

func calculaBMI(weight, height float64) float64 {
	return weight / (height * height)
}

func BMICalculator() {

	fmt.Println("How much do you weigh? (don't lie)")
	var weight float64
	fmt.Scan(&weight)

	fmt.Println("How tall are you? (barefoot)")
	var height float64
	fmt.Scan(&height)

	bmi := calculaBMI(weight, height)

	fmt.Printf("Right now your BMI is %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("You are underweight, add more potato to the broth")
	} else if bmi < 25 {
		fmt.Println("You have a normal weight, I have healthy envy of you")
	} else {
		fmt.Println("You are overweight, I know, the pandemic has affected us all")
	}
}
