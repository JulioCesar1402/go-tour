package main

import "fmt"

// ! funciona porem retorna um erro no codewars (out of range?)
// func Solution(word string) string {
// 	if len(word) == 1 {
// 		return word
// 	}
// 	length := len(word) - 1
// 	last_char := word[length:]
// 	word_without_last_char := Solution(word[0:length])
// 	return last_char + word_without_last_char
// }

func Solution(word string) (result string) {
	for _, i := range word {
		result = string(i) + result
	}
	return
}

func main() {
	solution := Solution("Julio Cesar")
	fmt.Println(solution)
}
