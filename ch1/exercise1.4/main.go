// Dup2 shows the couting and text from lines that appear more than once
// in the input. It reads from stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, "", counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts)
			f.Close()
		}
	}

	for line, filenames := range counts {
		fileCount := len(filenames)
		if fileCount == 1 {
			total := 0
			for _, count := range filenames {
				total += count
			}
			if total <= 1 {
				continue
			}
		}

		for name, count := range filenames {
			fmt.Printf("%d\t%s\t%s\n", count, line, name)
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(map[string]int)
		}
		counts[input.Text()][filename]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
