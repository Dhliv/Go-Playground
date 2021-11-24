package Essentials

import "fmt"

type figuras2D interface {
	area() float64
}

func calcularArea(f figuras2D) {
	fmt.Printf("El area de la figura %T es: %f\n", f, f.area())
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base, altura float64
}

func (c *cuadrado) area() float64 {
	return c.base * c.base
}

func (r *rectangulo) area() float64 {
	return r.altura * r.base
}

func Interfaz() {
	var c cuadrado = cuadrado{base: 10}
	var r rectangulo = rectangulo{base: 10, altura: 5}

	calcularArea(&c)
	calcularArea(&r)

	// Lista de interfaces (cualquier dato)
	l := []interface{}{"hola", false, 123321}
	fmt.Println(l...)
}
