// Echo3 prints the index and value of each of its
// command-line arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func printArgs(i int, args []string) {
	if len(args) == 0 {
		return
	}
	fmt.Printf("%d %s\n", i, args[0])
	printArgs(i+1, args[1:])
}

func main() {
	printArgs(0, os.Args[1:])
}
