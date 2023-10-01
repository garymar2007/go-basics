package main

import "fmt"

func main() {
	// var counter int32 = 0
	// for {
	// 	//rand.Seed(time.Now().UnixNano())  from g0 1.2 this is not required.
	// 	i := rand.Intn(100) + 1
	// 	counter++
	// 	if i == 99 {
	// 		break
	// 	}
	// }
	// fmt.Printf("Produce the number 99 in %d times", counter)

	//break with labels
	// label2:
	// for i := 0; i < 4; i++ {
	// 	//label1: //setup a label
	// 	for j := 0; j < 10; j++ {
	// 		if j == 2 {
	// 			//break label1
	// 			break label2
	// 		}
	// 		fmt.Println("j=", j)
	// 	}
	// }

	//Exe1: From 1 to 100, the first time when the sum of the numbers > 20
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("The current number is: ", i)
			break
		}
	}
	fmt.Println("The current sum is: ", sum)

	//Exe2: 3 login attemps, if username="gary" and password="123456", then login successfully.  Otherwise, prompt how many chances left.

}
