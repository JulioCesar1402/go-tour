package main

import (
	"bufio"
	"fmt"
	"os" // operation system
	// "strconv"
)

func main() {
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
