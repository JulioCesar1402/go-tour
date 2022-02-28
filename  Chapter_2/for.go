package main

import "fmt"

//	No golang, a estrutura de laço é apenas o "for"

func main() {
	// ? Existem duas formas de estruturar o for

	// * 1. passar valor inicial, condição de parada, ação a completar uma volta completa
	fmt.Println("========Case  1=========")
	sum := 0
	for i := 0; i < 10; i += 1 {
		sum += i
		fmt.Printf("Round: %v\n", i)
		fmt.Printf("Sum: %v\n", sum)
		fmt.Println("=-=-=-=-=-=-=-=-=-=-")
	}
	fmt.Printf("sum of all Numbers %v\n", sum)

	// * 2. condição de parada, também conhecido como "while" do Go
	sum_2 := 1

	for sum_2 < 1000 {
		sum_2 += sum_2
	}

	fmt.Println("========Case  2=========")
	fmt.Println(sum_2)

	/*
		& 	Bónus: existe o Forever:
		& 	for {
		& 	}
	*/
}