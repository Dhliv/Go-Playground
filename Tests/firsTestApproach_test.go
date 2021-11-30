package Tests

import (
	"testing"
)

// We got to execute _go test_ inside this folder, but be careful, we NEED a module of the proyect.
// If you dont export the functions that execute the tests, they will not get evaluated.

// ! In order to get the coverage of the testing, we NEED both files (the test and the file to test)
// ! inside the same directory, otherwise it wont work.

/*
	We use _go test -coverprofile=out_ to get something pretty weird.

	We use _go tool cover -func=out_ to see in a human way what the command up of this really means:
	it gives you the rates of coverage (in tests) that we have for every function that we have in the
	file that we are testing.

	We can also use the command _go tool cover -html=out_ to see all of the former command, but in a
	visual way inside our browser. Pretty cool.

	Additionally, we can use the command _go test -cpuprofile=cpufile_ to know how is the cpu perfomance
	by executing the functions that we define over the cases that we describe in testing.

	With the former file, we can execute the command _go tool pprof cpufile_ to see something weird.
	Inside here, we have acces to an interactive console. We have commands like _top_ to see the rates
	of consume of the cpu by the different functions taht we are testing.
	There is more. With the command _list {function_name}_ we can see perfomance in time line to line
	of the function that we did choose to analyze.
	Finally, we can use _web_ or _pdf_ to generate visual reports of the execution. We an exit this
	interactive console by writing _quit_.
*/
func TestSum(t *testing.T) {
	tables := []struct {
		a, b, n int
	}{
		{1, 2, 3},
		{45, 45, 90},
		{123, -123, 0},
	}

	for _, v := range tables {
		total := Sum(v.a, v.b)
		if total != v.n {
			t.Errorf("Sum was incorrect, got %d. Expected %d\n", total, v.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a, b, res int
	}{
		{4, 1, 4},
		{100, 1, 100},
		{1231, 321, 1231},
		{123, 321, 321},
	}

	for _, v := range tables {
		max := Max(v.a, v.b)

		if max != v.res {
			t.Errorf("Max was incorrect, got %d. Expected %d\n", max, v.res)
		}
	}
}

func TestFibo(t *testing.T) {
	tables := []struct {
		n   int
		res int64
	}{
		{1, 1},
		{8, 34},
		{20, 10946},
		//{30, 1346269},
		//{40, 165580141},
		//{45, 1836311903},
	}

	for _, v := range tables {
		fibo := Fibo(v.n)

		if fibo != v.res {
			t.Errorf("Fibonacci was incorrect, got %d. Expected %d\n", fibo, v.res)
		}
	}
}
