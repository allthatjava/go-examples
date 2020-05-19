package main

import (
	"fmt"
	"strconv"
)

type Person struct{
	firstName string
	lastName string
	city string
	gender string
	age int
}

// Value receiver
func (p Person) greet() string {
	return "Hello, my name is "+p.firstName+" "+p.lastName+". I am "+ strconv.Itoa(p.age)
}

// Value receiver, so it value won't change
func (p Person) changeAge() {
	p.age = 60
	fmt.Println("Inside of changeAge func I am ", p.age)
}

// Pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

func main(){

	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Toronto", gender: "F", age: 20}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1.age)
	fmt.Println(person1.greet())

	person1.changeAge()
	fmt.Println(person1.greet())

	person1.hasBirthday()
	fmt.Println(person1.greet())
}
