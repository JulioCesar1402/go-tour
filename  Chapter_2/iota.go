package main

import "fmt"

const (
	v = iota // 0
	_ = iota // 1
	x = iota // 2
	_ = iota // 3
	z = iota // 4
	/*
		? É possível realizar formulas, todos os valores a partir de agora
		? serão "* 10"
	*/
	a = iota * 10 // 50
	b // 60
	c // 70
)

/*
	Não é preciso colocar "iota" em todos, o compilador sub-entende
	const (
		v = iota
		_
		x
		_
		z
	)
*/

func main() {
	fmt.Println(v, x, z)
	fmt.Println(a, b, c)
}
