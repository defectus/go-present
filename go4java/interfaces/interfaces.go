package main

import "math/big"
import "fmt"

type Number interface {
	Number() float64
}
type BigNumber big.Float

func (b BigNumber) Number() float64 {
	t := big.Float(b)
	f, _ := (&t).Float64()
	return f
}

func main() {
	PrintNumber(BigNumber(big.Float{}))
}

func PrintNumber(n Number) {
	fmt.Printf("and the number is %f", n.Number())
}
