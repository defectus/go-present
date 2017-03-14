package main

import (
	"fmt"
)

type A struct{}
type B struct{}

func (a *A) Print() { fmt.Println("A") }
func (b B) Print()  { fmt.Println("B") }
func main() {
	(&A{}).Print()
	B{}.Print()
}
