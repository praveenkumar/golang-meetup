package main

import "fmt"

type Animal struct {
	kind string
	name string
	age int
}


func main() {
	var cat Animal
	var dog Animal

	// Cat specification
	cat.kind = "cat"
	cat.name = "Jenny"
	cat.age  = 2

	// Dog specification
	dog.kind = "dog"
	dog.name = "Rocky"
	dog.age  = 5

	fmt.Printf("Cat: %#v\n", cat)
	fmt.Printf("Dog: %v\n", dog)
}
