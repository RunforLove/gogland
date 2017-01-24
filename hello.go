package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var b bool = true
	fmt.Println(b)
	c := 10
	fmt.Println(c)
	var a int = 10
	fmt.Println(a)
	var d,e,f = 10, 11, 12
	fmt.Println(d,e,&f)
	d,e = e,d
	fmt.Println(d,e)
	const k string = "abc"
	fmt.Println(k)
	const(
		Red = "hong"
		Green = "lv"
	)
	fmt.Println(Green, Red)
}