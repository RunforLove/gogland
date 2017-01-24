package main

import "fmt"

func main () {
	var a int =21
	var b int = 10
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)

	fmt.Println(a)
	a++
	fmt.Println(a)
	fmt.Println(b)
	b--
	fmt.Println(b)

	fmt.Println(a == b)
	fmt.Println(a != b)

	fmt.Println(true && false)

	k1 := 1
	k2 := 3
	fmt.Println(k1&k2) // 相与运算
	fmt.Println(k1|k2)
	fmt.Println(k1^k2) // 异或运算
	fmt.Println(k1<<1)
	fmt.Println(k2>>2)

	// 取地址和指针变量
	println("*****************")
	var k3 int32
	var k4 float32
	var k5 *int
	fmt.Printf("%T,%T,%T,%D",k3,k4,k5,k5)
}