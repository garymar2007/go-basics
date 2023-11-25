package utils

import "fmt"

//Make this function public.
func Cal(n float64, m float64, op byte) float64 {
	var result float64
	switch op {
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
