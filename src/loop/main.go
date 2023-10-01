package main

func main() {

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Hello World", i+1)
	// }

	// j := 1
	// for ; j <= 10; j++ {
	// 	fmt.Println("Hello World1", j)
	// }

	// k := 1
	// for k <= 10 {
	// 	fmt.Println("Hello World2", k)
	// 	k++
	// }

	// i := 1
	// for {
	// 	fmt.Println("Hello World3", i)
	// 	i++
	// 	if i > 10 {
	// 		break
	// 	}
	// }

	// var str string = "Hello world! Welcome to Gary's world"
	// for index, val := range str {
	// 	fmt.Printf("Index: %d, value=%c\n", index, val)
	// }

	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("Index: %d, value=%c\n", i, str[i])
	// }

	// var str string = "Hello world! Welcome to Gary's world! 禁止转载"
	// str2 := []rune(str) // convert str to []rune
	// for i := 0; i < len(str2); i++ {
	// 	fmt.Printf("Index: %d, value=%c\n", i, str2[i])
	// }

	// for index, val := range str {
	// 	fmt.Printf("Index: %d, value=%c\n", index, val)
	// }

	//Exe 1: in [1--100] print the number of a number which is divisible by 9 and the sum of those numbers
	// var count int
	// var sum int = 0
	// for i := 1; i <= 100; i++ {
	// 	if i%9 == 0 {
	// 		count++
	// 		sum += i
	// 	}
	// }
	// fmt.Printf("The number of numbers which is multiple of 9 is %d and the sum of those numbers is %d", count, sum)

	//Exe 2: print 0 + 6 = 6, 1 + 5 = 6... the bounds of 6
	// i := 0
	// var j int = 6
	// for i <= j {
	// 	fmt.Printf("%d + %d = %d\n", i, j-i, j)
	// 	i++

	// }

	//while
	// i := 10
	// for {
	// 	if i >= 0 {
	// 		fmt.Println(i)
	// 	} else {
	// 		break
	// 	}
	// 	i--
	// }

	//do-while
	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	if i >= 10 {
	// 		break
	// 	}
	// }

	//Exe 1: there 3 classes and 5 students in each class. output the average of each class and all the classes;
	//Exe 2: the number of passed students.
	// var number_classes int = 3
	// var number_students int = 5
	// var sum_per_class float32 = 0.0
	// var sum_all_classes float32 = 0.0
	// var passCounter int = 0
	// for i := 1; i <= number_classes; i++ {
	// 	for j := 1; j <= number_students; j++ {
	// 		fmt.Printf("Please enter the score of class-%d, student-%d\n", i, j)
	// 		var score float32
	// 		fmt.Scanln(&score)
	// 		sum_per_class += score
	// 		if score >= 60 {
	// 			passCounter++
	// 		}
	// 	}
	// 	fmt.Printf("The average of class-%d is %v\n", i, sum_per_class/float32(number_students))
	// 	sum_all_classes += sum_per_class
	// 	sum_per_class = 0.0
	// }
	// fmt.Printf("The average of all classes is %v\n", sum_all_classes/float32(number_students*number_classes))
	// fmt.Printf("The number of students who passed this exam: %d", passCounter)

	//Exe 3: Print out Hollow Pyramid
	// fmt.Println("Please input the number of layers:")
	// var layers int
	// fmt.Scanln(&layers)
	// for i := 1; i <= layers; i++ {
	// 	for k := 0; k < layers-i; k++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for j := 1; j <= 2*i-1; j++ {
	// 		if j == 1 || j == 2*i-1 || i == layers {
	// 			fmt.Print("*")
	// 		} else {
	// 			fmt.Print(" ")
	// 		}

	// 	}
	// 	fmt.Println()
	// }

	//Exe 4: 9 X 9 multiplication map
	// for i := 1; i <= 9; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d x %d = %d\t", j, i, i*j)
	// 	}
	// 	fmt.Println()
	// }

}
