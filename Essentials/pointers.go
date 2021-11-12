package Essentials

import "fmt"

type str struct {
	left *str
}

func main4() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, b, *b)
	*b = 321 // they work just like in C++, NICE
	fmt.Println(a, b, *b)

	c := [...]int{1, 2, 3}
	d := &c[0]
	e := &c[1] // we cant do pointers arithmetic.

	fmt.Printf("C: %v, d: %v and e: %v\n", c, d, e)

	var t str
	t.left = new(str)
	fmt.Println(t.left.left) // we dont need to complicate ourselves with deferencing the pointer inside a struct.

	//we have to be careful when dealing with slices and maps, cause we cant make a copy of them, we are always referencing them.
}
