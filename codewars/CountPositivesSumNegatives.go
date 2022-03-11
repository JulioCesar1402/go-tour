package main

import "fmt"

/*
	Given an array of integers.

	Return an array, where the first element is the count of positives
	numbers and the second element is sum of negative numbers.
	0 is neither positive nor negative.

	If the input is an empty array or is null, return an empty array.
	Example

	For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15],
	you should return [10, -65].
*/

func CountPositivesSumNegatives(numbers []int) []int {
	var countOfPositives, sumNegatives int

	for _, i := range numbers {
		if i > 0 {
			countOfPositives += 1
		} else {
			sumNegatives += i
		}
	}

	res := []int{countOfPositives, sumNegatives}
	return res
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Println(CountPositivesSumNegatives(array))
}
