package main

import "fmt"

func main() {

	// for i := 0; i < 4; i++ {
	// 	//label1: //setup a label
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			continue
	// 		}

	// 		fmt.Println("j=", j)
	// 	}
	// }

	// here:
	// 	for i := 0; i < 4; i++ {
	// 		//label1: //setup a label
	// 		for j := 0; j < 10; j++ {
	// 			if j == 2 {
	// 				continue here
	// 			}

	// 			fmt.Println("j=", j)
	// 		}
	// 	}

	//exe 1: print out odd number in [1-100]
	// for i := 1; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//exe 2: input two number from keyboard and sum them until 0
	// var numOfPositive int
	// var numOfNegative int
	// var num int16
	// for {
	// 	fmt.Println("Please input a number:")
	// 	fmt.Scanln(&num)
	// 	if num == 0 {
	// 		break
	// 	}

	// 	if num > 0 {
	// 		numOfPositive++
	// 		continue
	// 	}

	// 	numOfNegative++
	// }

	// fmt.Printf("The number of positive number is %d, and the number of negative number is %d\n", numOfPositive, numOfNegative)

	//exe 3: someone has R100000, and he needs to pay the toll gate every time: if his money > 50000, 5% is required otherwise, R1000 is required.
	var numOfTollGate int
	var money int = 100000
	for {
		if money > 50000 {

			money -= 0.05 * 50000
		} else {
			money -= 1000
		}
		numOfTollGate++

		if money <= 0 {
			break
		}
	}
	fmt.Printf("This person has passed %d Toll Gates.\n", numOfTollGate)
}
