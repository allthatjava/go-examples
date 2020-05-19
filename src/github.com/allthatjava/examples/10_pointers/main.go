package main

import "fmt"

func main(){

	a := 5
	b := &a

	fmt.Println(a)		// value of a
	fmt.Println(b)		// value of b (memory address)
	fmt.Println(*b)		// Value of where b points to

	fmt.Println(*&a)	// value of what a's memory address has

}
