package main

import (
	"fmt"
)

func main() {
	var i int = 21
	var j bool = true
	ya := "Я"
	var f float64 = 123.456

	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	fmt.Printf("%v \n", j)
	fmt.Printf("%q \n", ya)
	fmt.Printf("%b \n", i)
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%X \n", 15)
	fmt.Printf("%+q \n", ya)
	fmt.Printf("%.3f \n", f)
	fmt.Printf("%f \n", f)
	fmt.Printf("%E \n", f)
}
