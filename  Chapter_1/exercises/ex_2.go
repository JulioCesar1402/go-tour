package main

import "fmt"

var x int
var y string
var z bool


func main()  {
	s := fmt.Sprintln(x, "-", y, "-", z)
	fmt.Println(s)
}