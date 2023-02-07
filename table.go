package main

import (
	"fmt"
)

func main() {

	var num int
	fmt.Println("enter num:")
	fmt.Scanln(&num)
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "X", i, "=", num*i)
	}
}
