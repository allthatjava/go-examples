package main

import "fmt"

func main(){

	// In Long Form -------------
	//// Init Map variable
	//emails := make(map[string]string)
	//
	//// Insert Key, Value
	//emails["Bob"] = "bob@email.com"
	//emails["Sharon"] = "sharon@email.com"
	//emails["Mike"] = "mike@email.come"

	// In Short form -----------
	emails := map[string]string{"Bob":"bob@email.com", "Sharon":"sharon@email.com", "Mike":"mike@email.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete
	delete(emails, "Bob")
	fmt.Println(emails)

}
