package CP

import (
	"bufio"
	"fmt"
	"os"
)

var rdr *bufio.Reader = bufio.NewReader(os.Stdin)
var wr *bufio.Writer = bufio.NewWriter(os.Stdout)

func Printf(f string, a ...interface{}) {
	fmt.Fprintf(wr, f, a...)
}

func Scanf(f string, a ...interface{}) {
	fmt.Fscanf(rdr, f, a...)
}

func Flush() {
	wr.Flush()
}

func Print(val interface{}) {
	fmt.Fprintf(wr, "%v\n", val)
}
