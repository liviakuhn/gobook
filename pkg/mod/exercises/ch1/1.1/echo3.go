// Echo3 prints its command-line arguments as well as the name
// of the command that invoked the echo program.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[0])
}
