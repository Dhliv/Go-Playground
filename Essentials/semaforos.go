package Essentials

import (
	"fmt"
	"sync"
	"time"
)

func Semaforos() {
	var wg sync.WaitGroup
	c := make(chan int, 2)

	for i := 0; i < 10; i++ {
		c <- 1 // With this, we limitate the subroutines created, and we wait till them finish their proccess.
		wg.Add(1)
		go takeTime(i, &wg, c)
	}

	wg.Wait()
}

func takeTime(v int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	fmt.Printf("Process %d started\n", v)
	time.Sleep(2 * time.Second)
	fmt.Printf("Process %d finished\n", v)
	<-c
}
