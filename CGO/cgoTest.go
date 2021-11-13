//main.go
package CGO

/*
  #include "hello.c"
  #include "sum.c"
*/
import "C"
import (
	"errors"
	"log"
)

//Hello is a C binding to the Hello World "C" program. As a Go user you could
//use now the Hello function transparently without knowing that it is calling
//a C function
func hello() error {
	_, err := C.Hello() //We ignore first result as it is a void function
	if err != nil {
		return errors.New("error calling Hello function: " + err.Error())
	}

	return nil
}

func CGOTest() {
	//Call to void function without params
	err := Hello()
	if err != nil {
		log.Fatal(err)
	}
}
