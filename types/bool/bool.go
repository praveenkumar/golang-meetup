// The Boolean type (bool) can be one of two values, either true or false

package main

import "fmt"

func main() {
	x := 32
	y := 60

	// Bool value with comparison operators
	fmt.Println("x == y:", x == y)
	fmt.Println("x != y:", x != y)
	fmt.Println("x < y:", x < y)
	fmt.Println("x > y:", x > y)
	fmt.Println("x <= y:", x <= y)
	fmt.Println("x >= y:", x >= y)

	// Change the control flow of program using bool
	if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is greater than y")
	}

	// String comparison (Case sensitive)
	Foo := "Foo"
	foo := "foo"
	fmt.Println("Foo == foo", Foo == foo)

	// logical Operator
	// && (x && y) is the and operator. It is true if both statements are true.
	// || (x || y) is the or operator. It is true if at least one statement is true.
	// ! (!x) is the not operator. It is true only if the statement is false.
	fmt.Println((17 > 15) && (9 < 8))
	fmt.Println((23 == 23) || (8 != 8))
	fmt.Println(!(12 <= 11))
}

