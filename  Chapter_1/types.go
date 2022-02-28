package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
	bool
	string
	! int sendo o mais comum para números inteiros
	int, int8, int16, int32, int64
	uint, uint8, uint16, uint32, uint64
	byte -> pseudónimo para uint8
	rune -> pseudónimo para uint32
	float32, float64
	complex64, complex128

*/

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt((-5 + 12i))
)

func types() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// assim como int, float recebe 0 como valor padrão
	var f float64
	fmt.Printf("%v\n", f)
}

func convertType() {
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x,y,z)
}

func main() {
	types()
	fmt.Println("======================")
	convertType()
	fmt.Println("======================")

}