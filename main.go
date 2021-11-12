package main

// This file is just to test all the things that Im learning.

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Dhliv/Go-Playground/Essentials"
)

var rdr *bufio.Reader = bufio.NewReader(os.Stdin)
var wr *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) {
	fmt.Fprintf(wr, f, a...)
}

func scanf(f string, a ...interface{}) {
	fmt.Fscanf(rdr, f, a...)
}

func main() {
	var s string
	scanf("%s\n", &s)
	printf("Lo que escribió fue: %s\n", s)

	wr.Flush()

	Essentials.Interfaz()
}
