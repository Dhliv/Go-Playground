package Essentials

import "fmt"

func write(text string, c chan<- string) { // we specify that the channel is of type in (it just receives data)
	c <- text // we insert text inside the channel c.
}

func print(c <-chan string) { // we specify that the channel is of type out (it just retrieves data)
	text := <-c
	fmt.Println(text)
}

func message(text string, c chan string) {
	c <- text
}

func Channels() {
	c := make(chan string, 1)
	c2 := make(chan string, 2)
	c2 <- "Mensaje1"
	c2 <- "Mensaje2"

	fmt.Println("Hello")
	go write("Bye", c)
	fmt.Println(<-c) // we get out the data inside the channel

	fmt.Println(len(c2), cap(c2))
	close(c2) // This will close the channel (it wont receive anymore data)

	// We got to close the channel before exploring it.
	for msg := range c2 {
		fmt.Println(msg)
	}

	email1 := make(chan string)
	email2 := make(chan string)
	go message("Mensaje1", email1)
	go message("Mensaje2", email2)
	for i := 0; i < 2; i++ { // We got to make sure of the number of channels to know what to do
		select { // we select the channel depending on the one that did receive the data. Pretty cool.
		case m1 := <-email1:
			fmt.Println("Email recibido de email1:", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2:", m2)
		}
	}
}
