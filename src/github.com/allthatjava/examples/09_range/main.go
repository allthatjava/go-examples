package main

import "fmt"

func main(){

	ids := []int{11,23,43,55,33,6}

	// print with 'index i'
	for i, id := range ids {
		fmt.Printf("%d - ID:%d\n", i, id)
	}

	// If you don't want to use 'index i'
	for _, id := range ids {
		fmt.Printf("ID:%d\n", id)
	}

	// Use case: Summing the ids
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Use case: Range Map
	emails := map[string]string{"Bob":"bob@email.com", "Sharon":"sharon@email.com", "Mike":"mike@email.com"}
	for k, v := range emails{
		fmt.Printf("%s: %s\n", k , v)
	}

	// Use case: Range Map with just Keys
	for k := range emails{
		fmt.Printf("%s\n", k )
	}
}
