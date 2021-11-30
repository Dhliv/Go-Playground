package Essentials

import (
	"fmt"
	"time"
)

func Multiplexacion() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go emulateSomething(d1, c1, 1)
	go emulateSomething(d2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-c1:
			fmt.Println(channelMsg1)

		case channelMsg2 := <-c2:
			fmt.Println(channelMsg2)
		}
	}
}

func emulateSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
