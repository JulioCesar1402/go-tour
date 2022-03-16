package main

import "fmt"

func main() {
	const (
		first = iota + 2023
		second
		third
		fourth
	)
	fmt.Printf("%v - %v - %v - %v\n", first, second, third, fourth)

}
