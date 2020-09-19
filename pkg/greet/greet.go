package greet

import (
	"fmt"
	"io"
)

//Greet the person
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello %s !\n", name)
}
