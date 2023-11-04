package main

import "fmt"

func calculator(n float64, m float64, operator byte) float64 {
	var result float64
	switch operator {
	case '+':
		result = n + m
	case '-':
		result = n - m
	case '*':
		result = n * m
	case '/':
		result = n / m
	default:
		fmt.Println("Invalid operator")

	}

	return result
}

func main() {
	fmt.Println("Please input two numbers and an operator:")
	var num1 float64
	var num2 float64
	var op byte

	fmt.Scanf("%f %f %v", &num1, &num2, &op)
	result := calculator(num1, num2, op)
	fmt.Println("result = ", result)
}
