package foobar

import (
	"fmt"
)

func Foobar() {
	fmt.Print("Ingrese un numero entero: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return
	}
	for i := 1; i <= input; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FooBar")
		} else if i%3 == 0 {
			fmt.Print("Foo")
		} else if i%5 == 0 {
			fmt.Print("Bar")
		} else {
			fmt.Print(i)
		}

		if i < input {
			fmt.Print("->")
		}
	}
	fmt.Println()
}
