package main

import "fmt"

func main() {
	// Increase their size dynamically
	// Since it can be expanded to add more elements as needed, more commonly used than arrays.

	a := []string{"Hello", "from", "golang", "meetup"}
	fmt.Printf("Type %T, Length %d, Cap %d \n", a, len(a), cap(a))

	// Create a slice from array
	arr := [6]int{10, 11, 12, 13, 14,15}
	myslice := arr[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	// append to slice
	myslice2 := arr[3:5]
	fmt.Printf("myslice2 = %v\n", myslice2)
	appendedSlice := append(myslice, myslice2...)
	fmt.Println(appendedSlice)
}
