package main

import "fmt"

func simple_if(x int) string {
	if x < 10 {
		return "Yeahh, it's greater than 10"
	}
	return "Sorry, it's lower than 10"
}

func if_with_short_declaration(x int) string {
	if x < 10 {
		return "Yeahh, it's greater than 10"
	}
	return "Sorry, it's lower than 10"
}

func main() {
	t := simple_if(1)
	q := simple_if(23)
	fmt.Println(t)
	fmt.Println(q)
}