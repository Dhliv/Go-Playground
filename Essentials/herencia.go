package Essentials

import "github.com/Dhliv/Go-Playground/CP"

type persona struct {
	age  int
	name string
}

type contrato struct {
	salario int
}

type employee struct {
	persona
	contrato
}

func (p *persona) getName() string {
	return p.name
}

func (p *persona) setName(name string) {
	p.name = name
}

func (c *contrato) getSalario() int {
	return c.salario
}

func (c *contrato) setSalario(salario int) {
	c.salario = salario
}

func Herencia() {
	defer CP.Flush()

	var e *employee = new(employee)
	e.setName("Reynell")
	e.setSalario(4500000)

	CP.Print(e)
	CP.Print(e.getName())
	CP.Print(e.getSalario())
}
