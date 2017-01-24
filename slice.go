package main

import "fmt"

func main () {
	var numbers = make([] int,3,5)
	printSlice(numbers)

	var k1 [] int
	printSlice(k1)
	if(k1 == nil) {
		fmt.Println("slice is empty")
	}

	k2 := [] int {0,1,2,3,4,5,6,7,8,9}
	printSlice(k2)
	fmt.Println("origin slice = ", k2)
	fmt.Println("slice 1~4 = ", k2[1:4]) //到4，而不包含索引4
	fmt.Println("slice ->3 = ", k2[:3])
	fmt.Println("slice 4-> = ", k2[4:])

	k3 := make([]int, 0, 5)
	printSlice(k3)

	k4 := k2[:2]
	printSlice(k4)

	k5 := k2[2:5]
	printSlice(k5)
}
func printSlice (num [] int) {
	fmt.Printf("len = %d; cap = %d; slice = %v\n",len(num),cap(num),num)
}