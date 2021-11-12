package RandomThings

import (
	"fmt"
	"reflect"
)

// this is brutal:
var (
	_                = 123 // this just means taht we are throwing away this result, this symbol itsn't used by itself.
	something string = "aksjdhaksdjhaksjdh"
	number    int    = 321231
)

type interval struct {
	left  int
	right int
}

type st struct {
	interval
	segt  *st
	m     int
	texto string `i can put anything that I want` // This is a tag
}

func main2() {
	var x int64            // var name and type.
	const a int = 42       // constants have to be assigned in compile time, not after that or we will get an error.
	var Something int = 32 // with this name (first letter in uppercase, the variable gets exported)
	var something int = 32

	x = 34431
	grades := [3]int{1, 2, 3}   // A way to declare array
	Grades := [...]int{1, 2, 3} // another way to declare array
	var students [32]string
	students[2] = "test"

	y := int64(23121) // If we dont use a variable in compilation time, it gets detected as an error.
	fmt.Println(x + y + int64(Something) + int64(something) + int64(a))
	fmt.Printf("grades %v, Grades: %v and Students: %v\n", grades, Grades, students)

	//Creating a map:
	leMape := map[string]int{
		"something":      123,
		"algo":           1234,
		"something else": 453,
	}
	leMape["adicion"] = 321231
	delete(leMape, "algo")

	_, ok := leMape["algo"] // In this way we can know if the key its not present in the map.

	if _, ok := leMape["algo"]; ok { // the key was in the map.
		fmt.Println("This is the way to initialize something inside an if statement")
	} else if !ok {
		fmt.Println("Its not inside the map")
	} else {
		fmt.Println("it never gets here")
	}

	fmt.Println(leMape)
	fmt.Printf("Acces to map by []: %v\n", ok)

	t := st{
		m: 123,
	}
	t.left = 123     // we establish values from an external structure here. Nice
	t.segt = new(st) // We inicialize the pointer, cool

	var exOfSt st = st{
		interval: interval{left: 123},
		m:        123,
		segt:     new(st),
	}
	ty := reflect.TypeOf(st{})
	field, _ := ty.FieldByName("texto")
	fmt.Println(t, exOfSt)
	fmt.Println(field.Tag) // way to get the tag

	var n int = 10
	var i int = 0
	for i := 0; i < n; i++ { // this is ugly as fuck
		if i&1 == 1 {
			continue
		}
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	for i < n { // this is actually a while loop
		fmt.Printf("%v ", i)
		i++
	}
	fmt.Println()

	for { // while(true), lol
		break
	}

	for k, v := range Grades { // seems like a for each
		fmt.Printf("(Grades[%v] = %v)  ", k, v)
	}
}
