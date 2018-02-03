package main

import "fmt"

func p(a ...interface{}) {
	fmt.Println(a)
}
func main() {
	ex43()
	sec41()
}

func sec41() {
	fmt.Println("Fixed length sequence of zero or more elements of a particular type")
	var a [3]int
	fmt.Println(a[len(a)-1]) // 0

	if cap(a) != len(a) {
		fmt.Println(cap(a))
		fmt.Println(len(a))
		fmt.Println("Fixed length!!!!!")
	}

	fmt.Println("Print indices and elements:")
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Array literal syntax
	_ = [...]int{1, 2, 3}

	p([...]int{1} == [...]int{1})

}

func zeroCopy(a [32]byte) {
	return a
}

func zeroPtr(ptr *[32]byte) {
}

// Rewrite reverse to use an array pointer instead of a slice
func ex43() {

}
