package main

import "fmt"

func main() {
	var number int = 200
	fmt.Printf("%d - %b - %#x\n", number, number, number)

	var number_2 int = number << 1
	fmt.Printf("%d - %b - %#x\n", number_2, number_2, number_2)

}