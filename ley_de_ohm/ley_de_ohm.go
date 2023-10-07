package elyDeOhm

import (
	"fmt"
)

func LeyDeOhm() float32 {

	var v, i, r float32

	fmt.Print("Ingrese los numeros ( puede ser voltaje,resistencia y corriente) ")
	fmt.Scan(&v, &r, &i)
	// Comprueba si alguno de los argumentos es cero y calcula el valor faltante.
	if v > 0 && i > 0 && r > 0 {
		fmt.Print("Ingresaste todos los numeros, ya tienes todos los datos necesarios")
		return 0
	} else if v > 0 && r > 0 {
		return v / r
	} else if v > 0 && i > 0 {
		return v / i
	} else if r > 0 && i > 0 {
		return r * i
	}
	return 0
}
