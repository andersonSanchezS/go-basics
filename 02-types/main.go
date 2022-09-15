package main

import "fmt"

func main() {
	// go data types

	// bool
	var boolean bool = false
	// string
	var str string = "hello world"
	// int int8 int16 int32 int64
	var i int = 1
	var i8 int8 = 2
	var i16 int16 = 3
	var i32 int32 = 4
	var i64 int64 = 5
	// uint uint8 uint16 uint32 uint64 uintptr
	var ui uint = 6
	var ui8 uint8 = 7
	var ui16 uint16 = 8
	var ui32 uint32 = 9
	var ui64 uint64 = 10
	// byte - alias for uint8
	var b byte = 11
	// rune // alias for int32
	var r rune = 12
	// float32 float64
	var f32 float32 = 13.0
	var f64 float64 = 14.0
	// complex64 complex128
	var c64 complex64 = 15.0
	var c128 complex128 = 16.0

	fmt.Println(boolean, str, i, i8, i16, i32, i64, ui, ui8, ui16, ui32, ui64, b, r, f32, f64, c64, c128)

	// int, uint, uintptr are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
	// When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

	// zero values
	// The zero value is:
	// 0 for numeric types,
	// false for the boolean type, and
	// "" (the empty string) for strings.
}
