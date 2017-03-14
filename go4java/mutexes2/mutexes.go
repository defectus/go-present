package main

import (
	"fmt"
	"time"
)

func main() {
	fun := func(callSign int) {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Microsecond)
			fmt.Println("func ", callSign, " writing")
		}
	}
	go fun(1)
	go fun(2)
	time.Sleep(time.Millisecond)
}
