package main

import (
	"fmt"
)

func test() bool {
	fmt.Println("test....")
	return true
}

func main() {
	// var i int = 10
	// if i < 9 && test() {
	// 	fmt.Println("Ok...")
	// }

	// if i > 9 || test() {
	// 	fmt.Println("hello...")
	// }

	// var a int = 5
	// var b int = 10
	// fmt.Printf("before exchange: a=%d, b=%d\n", a, b)

	// a = a + b
	// b = a - b
	// a = a - b
	// fmt.Printf("after exchange: a=%d, b=%d\n", a, b)

	// var name string
	// var age byte
	// var salary float32
	// var isPass bool
	// fmt.Println("Please input name: ")
	// fmt.Scanln(&name)
	// fmt.Println("Please input age: ")
	// fmt.Scanln(&age)
	// fmt.Println("Please input salary: ")
	// fmt.Scanln(&salary)
	// fmt.Println("Please input whether he/she passes the test: ")
	// fmt.Scanln(&isPass)
	// fmt.Printf("Name: %v\n Age: %v\nSalary: %v\nPass the test? %v\n", name, age, salary, isPass)

	// fmt.Println("Please input name, age, salary, whether pass the test by seperating them with space:")
	// fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
	// fmt.Printf("Name: %v\n Age: %v\nSalary: %v\nPass the test? %v\n", name, age, salary, isPass)

	// fmt.Println(2 & 3)  //2
	// fmt.Println(2 | 3)  //3
	// fmt.Println(2 ^ 3)  //1
	// fmt.Println(-2 ^ 2) //-4
	// fmt.Println(1 >> 2) //0
	// fmt.Println(1 << 2) //4

	// if age := 20; age > 18 {
	// 	fmt.Println("You are an adult!")
	// }

	// var x int = 4
	// if x > 2
	// 	fmt.Println("Ok")

	//Ex 01
	//fmt.Println("Exericise 01")
	// var number1 int32
	// var number2 int32
	// var number3 float64
	// var number4 float64

	// fmt.Println("Please input an int32: ")
	// fmt.Scanln(&number1)

	// fmt.Println("Please input another int32: ")
	// fmt.Scanln(&number2)

	// if number1+number2 > 50 {
	// 	fmt.Println("The sum of two numbers is greater than 50!")
	// 	fmt.Println("Hello World")
	// } else {
	// 	fmt.Println("The sum of two numbers is less than 50")
	// }

	// fmt.Println("Please input a float64: ")
	// fmt.Scanln(&number3)

	// fmt.Println("Please input another float64: ")
	// fmt.Scanln(&number4)

	// if number3 > 10.0 && number4 < 20.0 {
	// 	fmt.Printf("The sum of these two float64: %v", number3+number4)
	// }

	// var number1 int32
	// var number2 int32

	// fmt.Println("Please input an int32: ")
	// fmt.Scanln(&number1)

	// fmt.Println("Please input another int32: ")
	// fmt.Scanln(&number2)

	// if (number1+number2)%3 == 0 && (number1+number2)%5 == 0 {
	// 	fmt.Println("The sum of these two int32 is divisiable by 3 and 5!")
	// }

	// var leapYear int32
	// fmt.Println("Please input a year: ")
	// fmt.Scanln(&leapYear)

	// if (leapYear%4 == 0 && leapYear%100 != 0) || leapYear%400 == 0 {
	// 	fmt.Printf("The year %d is a leap year", leapYear)
	// } else {
	// 	fmt.Println(leapYear, " is not a leap year")
	// }

	// var score float32
	// fmt.Println("Please input your score:")
	// fmt.Scanln(&score)

	// if score == 100.0 {
	// 	fmt.Println("You are awarded an BMW toy car as you achieved a score=", score)
	// } else if score > 80 && score <= 99 {
	// 	fmt.Println("You are awarded an iPhone15Plus as you achieved a score=", score)
	// } else if score >= 60.0 && score <= 80 {
	// 	fmt.Println("You are awarded an iPad as you achieved a score=", score)
	// } else {
	// 	fmt.Println("Sorry!  You have no any award!")
	// }

	// var isPass bool = true
	// if isPass = false; !isPass {
	// 	fmt.Printf("Ok!")
	// }

	// var a float64
	// var b float64
	// var c float64
	// var x1 float64
	// var x2 float64
	// fmt.Println("Please input the values of a, b, and c with space to seperate them:")
	// fmt.Scanf("%v %v %v", &a, &b, &c)

	// if b*b-4*a*c > 0 {
	// 	x1 = (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	// 	x2 = (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)
	// 	fmt.Printf("There are two roots for this: %v and %v", x1, x2)
	// } else if b*b-4*a*c == 0 {
	// 	x1 = -b / (2 * a)
	// 	fmt.Printf("There is only one root for this: %v", x1)
	// } else {
	// 	fmt.Printf("There is no root for this!")
	// }

	// var height int32
	// var money int64
	// var isHandsome bool
	// fmt.Println("Please input conditions of you: height money isHandsome")
	// fmt.Scanf("%v %v %v", &height, &money, &isHandsome)

	// if height >= 180 && money >= 10000000 && isHandsome {
	// 	fmt.Println("Must marry him!")
	// } else if height >= 180 || money >= 10000000 || isHandsome {
	// 	fmt.Println("Marry him anyway!")
	// } else {
	// 	fmt.Println("Not marry him")
	// }

	// var speed float32
	// var isMale bool
	// fmt.Println("Please input the speed and sex, with space to seperate them:")
	// fmt.Scanf("%v %v", &speed, &isMale)

	// if speed < 8 {
	// 	if isMale {
	// 		fmt.Println("You are enterning the final round in male group")
	// 	} else {
	// 		fmt.Println("You are enterning the final round in female group")
	// 	}
	// } else {
	// 	fmt.Println("You are not enterning the final round. Sorry!")
	// }

	var month int16
	var age int16
	fmt.Println("Please input the month and your age, with space to seperate them:")
	fmt.Scanf("%v %v", &month, &age)

	if month >= 4 && month <= 10 {
		if age < 18 {
			fmt.Println("The price is: Half price")
		} else if age > 60 {
			fmt.Println("The price is: a third price")
		} else {
			fmt.Println("The price is: 60")
		}
	} else {
		if age >= 18 && age <= 60 {
			fmt.Println("The price is: 40")
		} else {
			fmt.Println("The price is: 20")
		}
	}
}
