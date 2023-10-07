package everOrOdd

import (
	"fmt"
)

func EvenOrOdd() {
	var input int
	fmt.Print("Escriba un numero positivo: ")
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return
	}

	if input != 0 {
		if input%2 == 0 {
			fmt.Println("Es par")
			return
		} else {
			fmt.Println("Es impar")
			return
		}

	} else {
		fmt.Println("No es un numero calculable")
		return
	}
}
