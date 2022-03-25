package main

import "fmt"

func getUserAgeFromInput() int {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	return age
}

func getUserNameFromInput() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	return name
}

func getUserFavoriteColorFromInput() string {
	var color string
	fmt.Print("Enter your favorite color: ")
	fmt.Scan(&color)
	return color
}

func main() {
	name := getUserNameFromInput()
	age := getUserAgeFromInput()
	color := getUserFavoriteColorFromInput()

	fmt.Printf("%v - %v - %v\n", name, age, color)
}