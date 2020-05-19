package main

import (
	"fmt"
)

func main(){
	//var name string = "Brian"
	var age = 30

	// Shorthand
	name := "Brian"

	fmt.Println(name, age)		// Print variable's value
	fmt.Printf("%T %d\n", age, age)	// Print Type of the Var and value
}