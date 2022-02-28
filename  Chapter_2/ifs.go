package main

import "fmt"

// simple if
func isGreaterThenTen(x int) string {
	if x < 10 {
		return "Yeahh, it's greater than 10"
	}
	return "Sorry, it's lower than 10"
}

// if with short declaration
func positiveSubstract(x, y  int) int {
	if z := x - y; x > y {
		return z
	} else {
		return y - x
	}
}

func main() {
	t := isGreaterThenTen(1)
	q := isGreaterThenTen(23)
	fmt.Println(t)
	fmt.Println(q)

	a := positiveSubstract(2,3)
	b := positiveSubstract(45,9)
	fmt.Println(a)
	fmt.Println(b)
}