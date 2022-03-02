package main

import "fmt"

func withSliceOfBytes() {
	s := "hello world"
	sb := []byte(s)

	for _, v := range sb {
		fmt.Printf("%v - %T - %#U - %#X\n", v, v, v, v)
	}
}

func withoutSliceOfBytes() {
	s := "hello òĉ 世界"

	for _, v := range s {
		fmt.Printf("%b - %T - %#U - %#X\n", v, v, v, v)
	}
	fmt.Printf("\n=-=-=-=-=-=-=-=-=-=-=-=-=-=-=\n")
	for i := 0; i < len(s); i += 1 {
		fmt.Printf("%b - %T - %#U - %#X\n", s[i], s[i], s[i], s[i])
	}
}

func main()  {
	withSliceOfBytes()
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	withoutSliceOfBytes()
}