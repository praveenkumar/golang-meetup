package main

import (
	"fmt"
)

func main()  {
	// Raw string literal
	// character sequences between back quotes (back ticks).
	a := `Say "hello" to Go!`
	fmt.Println(a)

	// Multiple line string
	b := `This string is on 
multiple lines
within a single back 
quote on either side.`
	fmt.Println(b)

	// Interpreted String Literals
	// character sequences between double quotes
	c := "Say \"hello\" to Go!"
	fmt.Println(c)

	// String Concatenation
	d := "Hello " + "golang " + "meetup"
	fmt.Println(d)

	// Read string character by character using range
	for i, s := range d {
		fmt.Printf("%d : %s\n", i, string(s))
	}

	fmt.Println(len(d))
}
