package main

import "fmt"

func main(){

	// Long form of FOR
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++;
	}

	// Short form of FOR
	for i := 1; i<=5; i++{
		fmt.Println(i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++{
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}

	}

}