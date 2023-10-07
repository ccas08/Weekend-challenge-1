package main

import (
	"fmt"
	bmiCalculator "main/bmi_calculator"
	evenOrOdd "main/even_or_odd"
	foobar "main/foobar"
	leyDeOhm "main/ley_de_ohm"
	mario "main/mario"
	"os"
)

func main() {
	var inicio string

	fmt.Print("Ingrese nombre de la función a usar: ")
	args := os.Args
	_, err := fmt.Scan(&inicio)
	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return
	}
	args[0] = inicio

	switch args[0] {
	case "EvenOrOdd":
		evenOrOdd.EvenOrOdd()
	case "LeyDeOhm":
		fmt.Print(leyDeOhm.LeyDeOhm())
	case "Foobar":
		foobar.Foobar()
	case "BMICalculator":
		bmiCalculator.BMICalculator()
	case "Mario":
		mario.Mario()
	default:
		fmt.Println("No es una opción válida.")
	}
}
