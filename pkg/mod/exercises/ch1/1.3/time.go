// Time measures the difference in running time between different
// implementations of echo, the program that prints its
// command-line arguments as well as the name of the command that
// invoked it.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(os.Args[0])	
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(os.Args[0])
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[0])
}

func main() {
	start := time.Now()
	echo1()
	secsEcho1 := time.Since(start).Seconds()
	fmt.Printf("Runtime echo1: %.6f seconds\n", secsEcho1)
	echo2()
	secsEcho2 := time.Since(start).Seconds() - secsEcho1
	fmt.Printf("Runtime echo2: %.6f seconds\n", secsEcho2)
	echo3()
	secsEcho3 := time.Since(start).Seconds() - secsEcho1 - secsEcho2
	fmt.Printf("Runtime echo3: %.6f seconds\n", secsEcho3)
}
