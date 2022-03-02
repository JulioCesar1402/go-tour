package main

import "fmt"

var x bool

func main()  {
	fmt.Println(x) // zero value
	x = (10 > 3)
	fmt.Println(x) // true
	y := (23 < 7)
	z := (6 == 6)
	j := ( 2 != 3)
	fmt.Println(y) // false
	fmt.Println(z) // true
	fmt.Println(j) // true

}