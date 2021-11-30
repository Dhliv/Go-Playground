package Essentials

import "fmt"

func Pipelines() {
	gen := make(chan int)
	doubles := make(chan int)

	go generator(gen)
	go double(gen, doubles)
	printChannel(doubles)
}

func generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}

	close(c)
}

func double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}

	close(out)
}

func printChannel(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}
