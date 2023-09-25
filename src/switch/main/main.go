package main

import "fmt"

func test(b byte) byte {
	return b + 1
}

func main() {

	// var weekDay byte

	// fmt.Println("Please input from a-g:")
	// fmt.Scanf("%c", &weekDay)

	// switch weekDay {
	// case 'a':
	// 	fmt.Println("Monday")
	// case 'b':
	// 	fmt.Println("Tuesday")
	// case 'c':
	// 	fmt.Println("Wednesday")
	// case 'd':
	// 	fmt.Println("Thursday")
	// case 'e':
	// 	fmt.Println("Friday")
	// case 'f':
	// 	fmt.Println("Satursday")
	// case 'g':
	// 	fmt.Println("Sunday")
	// default:
	// 	fmt.Println("Unknown week day - wrong input")
	// }

	switch 'a' {
	case 'a':
		fmt.Println("Monday")
	default:
		fmt.Println("Unknown week day - wrong input")
	}

	switch test('a') {
	case 'b':
		fmt.Println("Tuesday")
	default:
		fmt.Println("Unknown week day - wrong input")
	}

	// var n1 int32 = 5
	// var n2 int32 = 10
	// switch n1 {
	// case n2, 10, 5:
	// 	fmt.Printf("Ok")
	// case 5:		//compilation error
	// 	fmt.Println("Not OK")
	// default:
	// 	fmt.Println("default!")
	// }

	//var age int = 30
	switch age := 40; {
	case age == 10:
		fmt.Println("age == 10")
	case age >= 20:
		fmt.Println("age >= 20")
	default:
		fmt.Println("No matched!")
	}

	var number int = 10
	switch number {
	case 10:
		fmt.Println("Ok1")
		fallthrough
	case 20:
		fmt.Println("Ok2")
	case 30:
		fmt.Println("Ok3")
	default:
		fmt.Println("Not matched!")
	}

	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("The type of x: %T", i)
	case int:
		fmt.Printf("x is int type")
	case float64:
		fmt.Printf("x is float64 type")
	case func(int) float64:
		fmt.Printf("x is func(int) type")
	case bool, string:
		fmt.Printf("x is bool or string type")
	default:
		fmt.Printf("x is unknown type")
	}
}
