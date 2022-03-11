package main

import "fmt"

/*
	You are given two interior angles (in degrees) of a triangle.

	Write a function to return the 3rd.

	Note: only positive integers will be tested.

	https://en.wikipedia.org/wiki/Triangle
*/

func OtherAngle(a int, b int) (thirdAngle int) {
	const sumOfTriangleAngle int = 180
	thirdAngle = sumOfTriangleAngle - (a + b)

	return
}

func main() {
	fmt.Println(OtherAngle(60, 30))
}
