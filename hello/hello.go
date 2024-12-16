package github.com/vitorfloriano/arundels-pogtools/hello

import (
	"io"
	"fmt"
)

func PrintTo(w io.Writer) {
	fmt.Fprintln(w, "Hello, world")
}	
