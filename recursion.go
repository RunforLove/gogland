package main

import "fmt"

	func Factorial(x int) (result int) {
		if x == 0 {
			result = 1;
		} else {
			result = x * Factorial(x - 1)
		}
		return ;
	}
	func main () {
		var i int = 15;
		fmt.Printf("%d的阶乘是 %d\n", i, Factorial(i))
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum)/float32(count)
	fmt.Printf("mean的值为:%f\n",mean)

}
