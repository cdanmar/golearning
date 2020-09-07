//Modify the echo program to print the index and value of each of its arguments, one per line.

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[1:] {
		s += sep + arg
		fmt.Println(i, s)
		sep = ""
	}
	fmt.Println(s)
}
