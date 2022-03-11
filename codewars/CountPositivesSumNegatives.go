package main

import "fmt"

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
