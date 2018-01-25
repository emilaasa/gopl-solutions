package main

import "fmt"

func main() {
	pointers()
	variables()
}

func variables() {
	fmt.Println("These functions are equivalent: ")
	fmt.Println(newInt())
	fmt.Println(newInt2())
}
func newInt() *int {
	return new(int)
}
func newInt2() *int {
	var dummy int
	return &dummy
}

func pointers() {
	x := 1
	p := &x         // & = get me the memory address of x
	fmt.Println(p)  // should print a memory address
	fmt.Println(*p) // should print the variable x
	*p = 2          // equivalent of saying x = 2
	fmt.Println(p)  // should print a memory address
	fmt.Println(*p) // should print the variable x
	x = 3
	fmt.Println(*p) // should print the variable x
}
