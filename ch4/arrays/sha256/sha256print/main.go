package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var n = flag.Bool("n", false, "use sha512")

func main() {
	fmt.Fprintf(os.Stdout, "%x\n", sha256.Sum256([]byte(os.Args[1])))
	if *n {
		fmt.Fprintf(os.Stdout, "%x\n", sha512.Sum512([]byte(os.Args[1])))
	}
}
