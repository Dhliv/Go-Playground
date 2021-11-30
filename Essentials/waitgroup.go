package Essentials

import (
	"sync"
	"time"

	"github.com/Dhliv/Go-Playground/CP"
)

func Waitgorup() {
	defer CP.Flush()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	wg.Wait() // This blocks the subroutine till all the processes added to it terminate they job.
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	CP.Printf("We started the process %d\n", i)
	time.Sleep(2 * time.Second)
	CP.Print("We finished the process")
}
