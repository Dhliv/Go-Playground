package Essentials

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(text)
}

func ConcurrencyFirstApproach() {
	var wg sync.WaitGroup

	wg.Add(3)
	go say("Hello", &wg)
	go say("world", &wg) // this executes the function concurrently

	go func(x int) {
		defer wg.Done()
		fmt.Println("El valor es:", x)
	}(46)

	wg.Wait()

	// time.Sleep(time.Second * 1)
}
