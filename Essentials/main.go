package main

import "fmt"

func main() {
	// Declaración de constantes
	const pi float64 = 3.141592
	const pi2 float64 = 3.1415

	fmt.Printf("Pi es: %f. Y un pi menos preciso es: %f\n", pi, pi2)

	// 3 formas de declarar variables
	base := 12
	var area int // Las variables vacías se inicializan a sus valores por defecto (zeros)
	var altura int = 10

	area = base * altura

	fmt.Printf("La base es: %d. La altura es: %d. El area es: %d\n", base, altura, area)
}
