package mario

import "fmt"

func Mario() {
	var height int

	fmt.Print("Pyramid height (1-8): ")
	_, err := fmt.Scan(&height)
	if err != nil {
		fmt.Println("Error al leer la altura:", err)
		return
	}

	if height <= 0 || height >= 9 {
		fmt.Println("La altura debe estar entre 1 y 8. Inténtalo de nuevo.")
		return
	}

	for i := 1; i <= height; i++ {
		// Imprime espacios en blanco para alinear a la derecha.
		for j := 0; j < height-i; j++ {
			fmt.Print(" ")
		}

		// Imprime hashes (#) para formar la pirámide.
		for k := 0; k < i; k++ {
			fmt.Print("#")
		}

		fmt.Println() // Nueva línea para la siguiente fila.
	}
}
