package main

import "fmt"

/*
	Consider an array/list of sheep where some sheep may be missing from their place.
	We need a function that counts the number of sheep present in the array (true means present).

	For example,

	[true,  true,  true,  false,
	  true,  true,  true,  true ,
	  true,  false, true,  false,
	  true,  false, false, true ,
	  true,  true,  true,  true ,
	  false, false, true,  true]

	The correct answer would be 17.

	Hint: Don't forget to check for bad values like null/undefined
*/

func CountSheeps(groupOfSheep []bool) (count int) {
	for _, sheep := range groupOfSheep {
		if sheep {
			count += 1
		}
	}
	return
}

func main() {
	sheep := []bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true}

	fmt.Println(CountSheeps(sheep))
}
