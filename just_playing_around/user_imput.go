package main

import (
	"bufio"
	"fmt"
	"os" // operation system
	"strconv"
)

func type_something() {
	// ? from os.Stdin (operation system) read the input from commandline
	scanner := bufio.NewScanner(os.Stdin)
	// ? this print is not needed to work but make more intuitive when the code is running
	fmt.Print("type something: ")
	// ? this actually caputere the line and storage in the scanner object
	scanner.Scan()
	// ? to get the storage input it's needed to put it into a variable
	input := scanner.Text()
	fmt.Printf("you typed: %q\n", input)
}

func get_year(message string) (input int64) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(message + ":\n> ")
	scanner.Scan()
	input, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	return
}

func guess_your_age() {
	actual_year := get_year("What year are we in")
	year_user_was_born := get_year("Type the year you were born")
	fmt.Printf("you will be %d years old at the and of %v\n", actual_year-year_user_was_born, actual_year)
}

func main() {
	// type_something()
	guess_your_age()
}
