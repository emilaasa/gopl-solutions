package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lineCounts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("lol", os.Stdin, lineCounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
				continue
			}
			countLines(arg, f, lineCounts)
			f.Close()
		}
	}
	for line, count := range lineCounts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}

}

func countLines(arg string, f *os.File, lineCounts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		lineCounts[arg+line] = lineCounts[line] + 1
	}
}
