package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	get_age := bufio.NewScanner(os.Stdin)
	fmt.Print("How old are you?\n> ")
	get_age.Scan()
	age := get_age.Text()
	fmt.Printf("Cool, you are %v years old\n", age)
}