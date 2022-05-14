package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// The predeclared architecture-independent numeric types
	//
	// uint8       the set of all unsigned  8-bit integers (0 to 255)
	var a uint8 = 255
	fmt.Printf("Type of a: %T, value: %d and size: %d byte\n", a, a, unsafe.Sizeof(a))
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	var b uint16 = 4
	fmt.Printf("Type of b: %T, value: %d and size: %d byte\n", b, b, unsafe.Sizeof(b))
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
	//
	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	//
	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers
	//
	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts
	//
	// byte        alias for uint8
	// rune        alias for int32
	fmt.Printf("Type of a: %T, value: %d and size: %d byte\n", a, a+1, unsafe.Sizeof(a))

	// uint     either 32 or 64 bits
	// int      same size as uint
	// uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
	var c int = 255
	fmt.Printf("Type of a: %T, value: %d and size: %d byte\n", c, c, unsafe.Sizeof(c))
}
