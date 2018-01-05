package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\t%v\n", os.Args[0], err)
			os.Exit(1)
		}
		// ex 1.7
		b, err := io.Copy(os.Stdout, res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: reading %s: %v\n", os.Args[0], url, err)
			os.Exit(1)
		}
		fmt.Println(b)
		fmt.Println(res.Status)
	}
}
