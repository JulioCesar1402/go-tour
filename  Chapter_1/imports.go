package main

import (
	"fmt"
	"math"
)

/*
	Or can be write like this

	import "fmt"
	import "math"
*/

func main() {
	res_1 := math.Sqrt(144)
	res_2 := math.Sqrt(-100)
	fmt.Printf("Now you have %g problems. \n", res_1)
	fmt.Printf("Now you have %g problems. \n", res_2)
}
