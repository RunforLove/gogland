package main

import "fmt"

func main () {
	var numbers []int
	printS(numbers)

	numbers = append(numbers,0)
	printS(numbers)

	numbers = append(numbers,1)
	printS(numbers)

	numbers = append(numbers, 2,3,4)
	printS(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers) * 2))

	copy(numbers1,numbers)
	printS(numbers1)
}

func printS (num [] int) {
	fmt.Printf("len = %d; cap = %d; slice = %v\n",len(num),cap(num),num)
}
