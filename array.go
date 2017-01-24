package main

import "fmt"

func main () {
	var n [10] int
	var i,j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n",j,n[j])
	}

	var k1 = 100;
	fmt.Printf("parameter is %x\n",&k1)

	var k3 int = 20
	var k4 *int
	k4 = &k3
	fmt.Printf("a's address is %x\n", &k3)
	fmt.Printf("a's address is %x\n", k4)
	fmt.Printf("a's value is %d\n", *k4)
}
