package main

import (
	"fmt"
	"os"
)

func main() {

	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	var name = "dolly"
	fmt.Fprint(os.Stdout, "hello, %s", name)
	fmt.Println(c)
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c = *c + ByteCounter(len(p))
	return len(p), nil
}
