package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cnt := make(map[string]int)
	//https://www.wikiwand.com/en/Standard_streams
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		cnt[line] = cnt[line] + 1
	}

	for line, n := range cnt {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
