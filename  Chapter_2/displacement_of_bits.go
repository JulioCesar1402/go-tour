package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB
	GB
	TB
)

func main() {
	fmt.Printf("binary\t\t\t\t\t\tdecimal\n")
	fmt.Printf("%b\t\t\t\t\t%d\n", KB, KB)
	fmt.Printf("%b\t\t\t\t%d\n", MB, MB)
	fmt.Printf("%b\t\t\t%d\n", GB, GB)
	fmt.Printf("%b\t%d\n", TB, TB)
}
