package main

// global variable declaration
var c, java, python bool = true, false, true

// Pi global constant variables declaration
const Pi = 3.14

func main() {
	// go variable types
	// bool
	// string
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte // alias for uint8
	// rune // alias for int32
	// float32 float64
	// complex64 complex128
	println(Pi)
	// type conversions
	bhaskar := "hero"
	rank := 1
	height := 5.11

	println(bhaskar, float32(rank), int(height))

	// this is a simple example demonstrating the declaration of variable in var
	var x int = 10
	println(x, c, java, python)
	// this is an example demonstarting the declaring and initializing of variables
	a, b := swappingOfNumbers(10, 20)
	println(a, b)
	//variables with initilizers so that we can omit the type
	var i, j, k = 1, 2, 3
	println(i, j, k)
}

// function demonstrating how the return works in go
func additionOfNumbers(x int, y int) int {
	return x + y
}

// function demonstrating how the return can handle two or more than two values
func swappingOfNumbers(x int, y int) (int, int) {
	return y, x
}

// function demonstrating how the naked return works in go
func nakedReturn(x int, y int) (a int, b int) {
	a = y
	b = x
	return
}
