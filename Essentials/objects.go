package Essentials

import (
	"github.com/Dhliv/Go-Playground/CP"
)

type empleado struct {
	id              int
	name, ocupacion string
}

func newEmpleado(id int, name, ocupacion string) *empleado {
	return &empleado{
		id:        id,
		name:      name,
		ocupacion: ocupacion,
	}
}

func Objects() {
	defer CP.Flush()
	var e *empleado = new(empleado) // Esto inicializa los objetos a su valor zero. Retorna un puntero.
	CP.Printf("%v\n", *e)

	var e2 *empleado = newEmpleado(1, "Alejandro", "Jardinero")
	CP.Printf("%v\n", *e2)
}
