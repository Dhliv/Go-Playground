package Essentials

import "fmt"

func prnt(msg *string, aux int) { // Funcs by defect pass copies of the data.
	*msg = "David Henao"
	fmt.Println(*msg, aux)
}

func multInputs(values ...int) { // If we want this with some more parameters, the ...type must be at the end.
	// Also, ...int create a slice, so be careful.
	fmt.Println(values)
	var sum int = 0

	for _, v := range values {
		sum += v
	}

	fmt.Println(sum)
}

func multType(n, a, m, e string, abc int) int {
	fmt.Println(n, a, m, e, " cool ", abc)
	return abc * abc
}

func sum(values ...int) (resul int) { // We declare a name for te return value, and its a variable. Cool
	for _, v := range values {
		resul += v
	}
	return
}

func div(a, b float64) (float64, error) { // we can return multiple things
	if b == 0.0 {
		return 0.0, fmt.Errorf("detected division by zero")
	}
	return a / b, nil // not error
}

// ! Methods:

type person struct {
	name     string
	apellido string
}

func (p person) greet() int { // This is a method, and its owner its person struct.
	// The parameter p is a copy of the original data, so be careful
	fmt.Println("Hola", p.name, p.apellido)
	return 0
}

func mainn() { // Same as variables in name convention, first letter in upper means its public, otherwhise is private
	var name string = "David"
	prnt(&name, 123) // pass the position of memory.
	a := multType("D", "av", "i", "d", 12)
	fmt.Println(a)
	multInputs(1, 2, 3, 4, 5, 6, 7, 8, 9)
	multInputs(1, 2, 3, 4, 5)

	abc := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(abc)

	d, err := div(5.0, 1.000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	var lambda func() = func() { // The best practice for this is to provide the paramaters that are used.
		// Is like that becasue with concurrency the things could have an odd behaviour.
		fmt.Println("Lambda")
	}

	var lambda2 func(int, int) int = func(a, b int) int {
		return a + b
	}

	lambda()
	lambda2(2, 4)

	var p person = person{name: "David", apellido: "Henao Martinez"}
	fmt.Println(p.greet())
}
