package main

import "fmt"

type coffee int
var x coffee

func main()  {
	fmt.Printf("%v - %T\n", x, x)
	x = 42
	fmt.Printf("%v - %T\n", x, x)
}