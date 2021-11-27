package Tests

import (
	"testing"

	"github.com/Dhliv/Go-Playground/Essentials"
)

func TestGetEmployeeByName(t *testing.T) {
	table := []struct {
		name             string
		mockFunc         func() // Its a function that emulates the returns of the original functions.
		expectedEmployee Essentials.Employee
	}{
		{
			name: "Fernando",
			mockFunc: func() {
				// We Override the original functions that the function that we gonna test use.
				// In this case, we Override 2 functions.
				// All the functions should be equal in name, parameters and returns.
				// We just should focus on the returns, 'cause at the end, we dont really prove it
				// in a real environment, we just place here what we are sure it will return to analyze
				// the results that the main function returns.
				Essentials.GetPersonaByName = func(name string) (Essentials.Persona, error) {
					return Essentials.Persona{
						Name: name,
						Age:  24,
					}, nil
				}

				Essentials.GetContrato = func() (Essentials.Contrato, error) {
					return Essentials.Contrato{
						Salario: 1200000,
					}, nil
				}
			},
			expectedEmployee: Essentials.Employee{
				Persona: Essentials.Persona{
					Age:  24,
					Name: "Fernando",
				},
				Contrato: Essentials.Contrato{
					Salario: 1200000,
				},
			},
		},
	}

	// We got to make sure that we save all the methods that we did override, so we can restore them.
	originalGetPersonaByName := Essentials.GetPersonaByName
	originalGetContrato := Essentials.GetContrato

	for _, test := range table {
		test.mockFunc()

		e, err := Essentials.GetEmployeeByName(test.name)

		if err != nil {
			t.Errorf("Error while getting Employee")
		}

		if e.Age != test.expectedEmployee.Age {
			t.Errorf("Error. Expected %v. Got instead %v", test.expectedEmployee.Age, e.Age)
		}

		if e.Name != test.expectedEmployee.Name {
			t.Errorf("Error. Expected %v. Got instead %v", test.expectedEmployee.Name, e.Name)
		}

		if e.Salario != test.expectedEmployee.Salario {
			t.Errorf("Error. Expected %v. Got instead %v", test.expectedEmployee.Salario, e.Salario)
		}

		Essentials.GetContrato = originalGetContrato
		Essentials.GetPersonaByName = originalGetPersonaByName
	}
}
