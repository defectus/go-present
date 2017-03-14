package main

import (
	"fmt"
)

func main() {
	array := [10]int{}
	fmt.Println("array length", len(array))
	slice := []int{1}
	fmt.Println("slice length", len(slice))
	slice = append(slice, []int{1, 2, 3}...)
	fmt.Println("slice extended length", len(slice))
	slice = slice [1:2]
	fmt.Println("slice shrunken length", len(slice))
}
