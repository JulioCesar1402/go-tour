package main

import "fmt"

func simpleIf(x int) string {
	if x < 10 {
		return "Yeahh, it's greater than 10"
	}
	return "Sorry, it's lower than 10"
}

func ifWithShortDeclaration(x int) string {
	if x < 10 {
		return "Yeahh, it's greater than 10"
	}
	return "Sorry, it's lower than 10"
}

func main() {
	t := simpleIf(1)
	q := simpleIf(23)
	fmt.Println(t)
	fmt.Println(q)
}