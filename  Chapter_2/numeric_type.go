package main

import (
	"fmt"
)

func main() {
	/*
		? uint8 e int8 se diferenciam pelo alcance de números que podem armazenar
		! var a uint8 = -1, retorna um erro pois vai de 0 a 255
		! fmt.Println(a)

		* var a int8 = -1, não retorna um erro pois vai de -128 a 128
		* fmt.Println(a)

		& essa diferença vale para todos os tipos de int(rune) e uint(byte), mudando apenas
		& a capacidade de armazenamento
	*/

	a, b, c := "e", "é", "世界"
	fmt.Printf("%v - %v - %v\n", a, b, c)

	d, e, f := []byte(a), []byte(b), []byte(c)
	fmt.Printf("%v - %v - %v\n", d, e, f)

	/*
		& Bonus Overflow:
		& Os tipos numericos possuem um limite, por exemplo
		& ao usar um uint16 (vai de 0 até 65535), "65535" é seu número limite
		& se tentar "var n uint16 = 65535" retornara um erro avisando que o
		& range de uint16 foi ultrapassado, por fim
		& se tentar "var n uint16 = 65535" e depois "n += 1", será retornado o valor "0"
		& isso acontece por seu limite ter sido alcançado, por não saber como lidar
		& com isso, a maquina irá interpretar que após "65535", virá o "0"
	*/
}
