package main

import "fmt"

func add(num1, num2 int) (result int) {
	result = num1 + num2
	return
}

/*
	Diferente dos parâmetros, em que o tipo pode ser omitido
	e então colocado no final, o return deve ter a quantidade tipos
	equivalente ao que é retornado
*/

func swap(x, y string) (string, string) {
	return y, x
}

/*
	Indicar o que a função retorna, chamado de retorno limpo,
	seu uso é recomendado em funções curtas, usado para fins de documentação.
	Em funções longas pode acabar prejudicando a legibilidade
*/

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	sum := add(10, 2)
	fmt.Println(sum)

	a, b := swap("Julio", "Cesar")
	fmt.Println(a, b)

	num_1, num_2 := split(35)
	fmt.Println(num_1, num_2)

}