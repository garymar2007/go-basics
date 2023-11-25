package main

import (
	"fmt"
	"goproject/src/project/utils"
)

func main() {
	fmt.Println("Please input two numbers and an operator:")
	var num1 float64
	var num2 float64
	var op byte

	fmt.Scanf("%f %f %c", &num1, &num2, &op)
	result := utils.Cal(num1, num2, op)
	fmt.Println("result ~= ", result)
}
