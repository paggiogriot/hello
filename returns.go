package main

import (
	"fmt"
)

func names() (string, string) {
	return "Foo", "Bar"
}

func main() {
	n1, n2 := names()
	fmt.Println(n1, n2)

	n3, _ := names()
	fmt.Println(n3)
}
