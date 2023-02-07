package main

import "fmt"

func main() {

	s1 := "abbccd"
	s2 := "abcbdd"
	var distance int
	fmt.Println(s1)
	fmt.Println(s2)
	r1 := []rune(s1)
	r2 := []rune(s2)
	fmt.Println(r1)
	fmt.Println(r2)
	for i, v := range r1 {
		if r2[i] != v {
			distance += 1
		}
	}
	fmt.Println(distance)
}
