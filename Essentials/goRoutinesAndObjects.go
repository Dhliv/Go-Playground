package Essentials

import (
	"fmt"
)

type objectSync interface {
	execute(chan string)
}

type st struct {
	l, r, m int32
}

func (st *st) st(l, r, m int32) {
	st.l = l
	st.r = r
	st.m = m
}

func (st *st) print() {
	fmt.Println(st.l, st.m, st.r)
}

func (st *st) execute(c chan string) {
	st.print()
	c <- "Done"
}

func execute(object objectSync) {
	c := make(chan string, 1)
	defer close(c)
	go object.execute(c)
	fmt.Println(<-c)
}

func GoRoutinesAndObjects() {
	var tr st
	tr.st(1, 100, 50)

	execute(&tr)
}
