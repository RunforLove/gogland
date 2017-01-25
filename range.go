package main

import "fmt"

func main() {
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums { // _省略了索引
		sum += num
	}
	fmt.Printf("sum = %d\n", sum)

	for i,num := range nums {
		if num == 3 {
			fmt.Println("current index is ", i)
		}
	}

	kvs := map[string]string{"a":"Apple","b":"Banana"}
	for k,v := range kvs {
		fmt.Printf("%s --> %s\n", k, v)
	}

	// 枚举Unicode字符串
	for i,c := range "go" {
		fmt.Println(i,c)
	}
}