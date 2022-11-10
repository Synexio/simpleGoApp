package main

import "fmt"

func main() {
	fmt.Println("Hello world !")

	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	fmt.Printf("len = %d cap = %d %v", len(s), cap(s), s)
}
