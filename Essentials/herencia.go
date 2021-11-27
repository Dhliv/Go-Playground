package Essentials

import (
	"fmt"
	"time"

	"github.com/Dhliv/Go-Playground/CP"
)

type Persona struct {
	Age  int
	Name string
}

type Contrato struct {
	Salario int
}

type Employee struct {
	Persona
	Contrato
}

func (p *Persona) getName() string {
	return p.Name
}

func (p *Persona) setName(name string) {
	p.Name = name
}

func (c *Contrato) getSalario() int {
	return c.Salario
}

func (c *Contrato) setSalario(salario int) {
	c.Salario = salario
}

var GetPersonaByName = func(name string) (Persona, error) {
	time.Sleep(time.Second * 2)
	// We simulate here a query to DB.
	return Persona{Name: name, Age: 24}, nil
}

var GetContrato = func() (Contrato, error) {
	time.Sleep(time.Second * 2)
	// We simulate here a query to DB.
	return Contrato{Salario: 1200000}, nil
}

func GetEmployeeByName(name string) (Employee, error) {
	var emp Employee
	p, err := GetPersonaByName(name)
	if err != nil {
		fmt.Errorf("This person dont exists")
		return emp, err
	}
	emp.Persona = p

	c, err := GetContrato()
	if err != nil {
		fmt.Errorf("Este Contrato no existe")
		return emp, err
	}
	emp.Contrato = c

	return emp, nil
}

func Herencia() {
	defer CP.Flush()

	var e *Employee = new(Employee)
	e.setName("Reynell")
	e.setSalario(4500000)

	CP.Print(e)
	CP.Print(e.getName())
	CP.Print(e.getSalario())
}
