package hello

import (
	"io"
	"fmt"
)

func PrintTo(w io.Writer) {
	fmt.Fprintln(w, "Hello, world")
}	
