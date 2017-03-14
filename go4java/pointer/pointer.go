package main

import (
	"fmt"
)

func main() {
	var a int
    var b *int
	a = 1
    b = &a
	fmt.Println("a is ", a, " and b is ", b, " which is ", *b)
}
