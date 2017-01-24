package main

import "fmt"

func main () {
	var k1 * int
	fmt.Printf("k1's value is %x\n",k1)
	fmt.Printf("k1's value is %x\n",&k1)
	fmt.Println(k1 == nil)
	fmt.Println(k1 != nil)
}
