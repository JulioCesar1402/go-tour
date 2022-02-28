package main

import "fmt"

// se não passado o valor, retorna false
var c, python bool

// assim como nos parametros, é possível omitir o tipo até que chegue no final
var h, j int = 1, 2

/*
	O tipo pode ser omitido quando tipos diferentes são atribuidos
	porém é considerado uma variavel "untyped"
*/
var javascript, java = true, "no!"

// name := "Julio Cesar"


func main() {
	// se não passado o valor, retorna 0
	var i int
	// se não passado o valor, retorna ""
	var s string
	fmt.Println(i, s)

	fmt.Println(c, python)

	fmt.Println(h, j)

	fmt.Println(javascript, java)

	/*
		É possível omitir "var" e o tipo, com a seguinte estrutura,
		porém só é utilizada dentro de uma função
		retornando:
		"syntax error: non-declaration statement outside function body"
	*/
	name := "Julio Cesar"
	fmt.Println(name)
}