package main

import "fmt"

func main(){

	// ARRAY ---
	//var fruitArr [2]string
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"
	// Or
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Slice ---
	fruitSlice := []string{"Apple", "Orange", "Grape","Pineapple"}
	fmt.Println(len(fruitSlice))		// Size of Slice
	fmt.Println(fruitSlice[1:3])		// Range
	fmt.Println(fruitSlice[2])


}