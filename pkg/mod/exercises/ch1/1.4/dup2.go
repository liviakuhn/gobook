// Dup2 prints the text and total count of lines that appear more than
// once in the input, as well as the number of times each duplicated
// line occurs in each of the files it read from, if any. It can also
// read from stdin.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int) // line -> filename -> count
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("", os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(arg, f, counts)
			f.Close()
		}
	}

	for line, fileCounts := range counts {
		total := 0
		for _, count := range fileCounts {
			total += count
		}
		if total > 1 {
			fmt.Printf("%s appears %d times in total:", line, total)
			for file, count := range fileCounts {
				// If file == "", we know the input was read from
				// stdin and the line appeared at least twice.
				if (file == "") {
					fmt.Printf("\n- %d times in stdin", count)
				} else if (count == 1) {
					fmt.Printf("\n- %d time in %s", count, file)
				} else {
					fmt.Printf("\n- %d times in %s", count, file)
				}
			}
			fmt.Println("\n")
		}
	}
}

func countLines(arg string, f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		// Ignore empty lines in the input.
		if (line != "") {
			if counts[line] == nil {
				counts[line] = make(map[string]int)
			}
			counts[line][arg]++
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
