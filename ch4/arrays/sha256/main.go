package main

import (
	"crypto/sha256"
	"fmt"
)

var a [3]int

func main() {
	fmt.Println(a[0])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d \n", v)
	}

	fmt.Printf("%T\n", a)

	r := [...]int{40: -1}
	fmt.Printf("%T\n", r)

	sha256Compare()

}
func sha256Compare() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	// Notes on Printf format string:
	// %x prints all the elements of an array or slice in hex
	// %t shows boolean
	// %T shows type
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(bitCompare(c1, c2))
}

func bitCompare(a1, a2 [32]byte) int {
	var count int

	for i, v := range a1 {
		if v != a2[i] {
			count++
		}
	}
	return count
}
