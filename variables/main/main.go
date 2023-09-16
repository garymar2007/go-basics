package main

import (
	"fmt"
	//"go/goproject/variables/model"
	"strconv"
	"unsafe"
)

var (
	d = 300
	e = 900
	f = "abc"
)

func main() {
	var i int
	fmt.Println("i=", i)

	var num = 10.11
	fmt.Println("num=", num)

	num1 := "tom"
	fmt.Println("name=", num1)

	a, b, c := 1, "gary", 15.09
	fmt.Println("a=", a, "b=", b, "c=", c)

	fmt.Println("d=", d, "e=", e, "f=", f)

	var g int64 = 900
	fmt.Printf("Type of g variable %T  the size of g is %d ", g, unsafe.Sizeof(g))

	var h = 1.1
	fmt.Printf("The default type of h is %T", h)

	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1=", c1)            //97
	fmt.Println("c2=", c2)            //48
	fmt.Printf("c1=%c c2=%c", c1, c2) // c1=a c2=0

	var c3 int = 22269
	fmt.Printf("c3=%c", c3) //国

	var str1 = "hello " + "world " + " how " + "are " + "you " +
		"I'm " + "fine " + "too"
	fmt.Println(str1)

	var n1 int32 = 100
	var n2 int64
	var n3 int8
	// n2 = n1 + 30		//error: int32 --> int64
	// n3 = n1 +20		//error: int32 --> int8
	n2 = int64(n1) + 30
	n3 = int8(n1) + 20
	fmt.Printf("n2=", n2, "n3=", n3)

	var nu1 int = 99
	var nu2 float64 = 23.456
	var nu3 int64 = 45687
	var bl bool = true
	var myChar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", nu1)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%f", nu2)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%t", bl)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(bl)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = strconv.FormatInt(int64(nu1), 10)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = strconv.FormatFloat(nu2, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = strconv.Itoa(int(nu3))
	fmt.Printf("str type %T str=%q\n", str, str)

	var s string = "true"
	var bo bool
	bo, _ = strconv.ParseBool(s) //this function returns two values (value bool, err error).  "_" means the value is ignored.
	fmt.Printf("b type %T b=%v\n", bo, bo)

	var str2 string = "1234567"
	var number int64
	var number2 int
	number, _ = strconv.ParseInt(str2, 10, 64)
	number2 = int(number)
	fmt.Printf("number type %T number=%v\n", number, number)
	fmt.Printf("number2 type %T number2=%v\n", number2, number2)

	var str3 string = "hello"
	var number3 int64 = 11
	number3, _ = strconv.ParseInt(str3, 10, 64)
	fmt.Printf("number3 type %T n3=%v\n", number3, number3)

	var ptr *int
	var number4 int = 999
	ptr = &number4
	fmt.Printf("ptr type %T ptr=%v *ptr=%v\n", ptr, ptr, *ptr)

	var number5 int = 300
	var number6 int = 400
	fmt.Printf("The original value of number5=%d and number6=%d\n", number5, number6)
	var ptr1 *int
	ptr1 = &number5
	*ptr1 = 100
	ptr1 = &number6
	*ptr1 = 200
	fmt.Printf("The changed value of number5=%d, number6=%d and ptr1=%v, *ptr1=%d\n", number5, number6, ptr1, *ptr1)

	//fmt.Printf("Today, I will introduce our hero -- %s", model.HeroName)

	fmt.Println(10 / 4) //output: 2
	var number7 float32 = 10 / 4
	fmt.Println(number7) //output: 2
	var number8 float32 = 10.0 / 4
	fmt.Println(number8) //output: 2.5

	var days int = 97
	var week int = days / 7
	var day int = days % 7
	fmt.Printf("%d weeks and %d days\n", week, day)

	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("%v F = %v C\n", huashi, sheshi)
}
