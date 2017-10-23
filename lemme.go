package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var python, java bool = true, false
// k :=3; // - not allowed outside func definition



func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func greeteng(firstName, lastName string) (string, string, string) {
	return "this is test", firstName, lastName
}

func printTypes() {
	// types
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	// print types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func displayEmpty() {
	// example of empty values
	var i int
	var f float64
	var b bool
	var s string

	fmt.Println("this are empty values")
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func main() {
	fmt.Printf("The pi is %g", math.Pi);
	fmt.Println("")
	fmt.Println("the result is: ", add(123, 45))
	fmt.Println("the result is: ", sub(123, 45))
	fmt.Println(greeteng("John", "Doe"))


	a, b, c := greeteng("Lisa", "Suomi")
	fmt.Println(c, b, a)

	// variables
	var i int = -0
	t := 3;
	fmt.Println(i, c, python, java, t)

	printTypes()
	displayEmpty()

}
