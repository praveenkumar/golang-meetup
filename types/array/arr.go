package main

import "fmt"

func main()  {
	// An ordered sequence of elements.
	// The capacity is defined at creation time.
	// Once allocated its size, the size can no longer be changed. (immutable)
	// Because the size of an array is static, it means that it only allocates memory once.

	// Defined by declaring the size, then the data type with the values defined between curly brackets { }.

	a := [3]string{"hello", "from"}
	fmt.Printf("Type %T, Length %d\n", a, len(a))
	fmt.Println(a)

	a[2] = "meetup"
	fmt.Printf("Length %d\n", len(a))
	fmt.Println(a)

	// a[3] = "out of bound error"
	// fmt.Println(a)

	// range over array
	for i, j := range a {
		fmt.Printf("index: %d, value: %s\n", i, j)
	}
}
