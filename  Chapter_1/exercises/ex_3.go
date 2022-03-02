package main

import "fmt"

var x int = 42
var y string = "monts"
var z bool = true


func main()  {
	s := fmt.Sprintf("%v\n%v\n%v\n", x, y, z)
	fmt.Println(s)
}