package RandomThings

import (
	"fmt"
	"log"
)

func panicking() {
	defer func() {
		if err := recover(); err != nil {
			// we cant invoke recover wiouth defer, cause the behaviour of panic will cause that we dont catcht the exceptions
			if err == "DIVISION BY 0 IS NOT ALLOWED" {
				fmt.Println("FOUND DIVISION BY 0") // we handle the exception that can trigger.
			} else { // we print an error log in case that we cant handle it, and make a panic alert to make the program stops.
				log.Println("Error: ", err)
				panic("CANT HANDLE THE POWER")
			}
		}
	}()

	b, c := 1, 0
	var alert string = "DIVISION BY 0 IS NOT ALLOWED"
	if c == 0 {
		panic(alert) // It just braks down the app, and its ecxecuted after the defer actions
	} else {
		fmt.Println(b / c)
	}
}

func main3() {
	//defer executes the things after the function ends its execution, and it works by LIFO.
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")

	a := "algo"
	defer fmt.Println(a) // defer saves the changes until its executed, but dont consider changes after call it.
	a = "change"
	panicking()
}
