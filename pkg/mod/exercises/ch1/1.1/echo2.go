// Echo2 prints its command-line arguments as well as the name
// of the command that invoked the echo program.

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(os.Args[0])
}
