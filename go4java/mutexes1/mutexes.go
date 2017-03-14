package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mutex := sync.Mutex{}
	fun := func(callSign int) {
		mutex.Lock()
		for i := 0; i < 5; i++ {
			time.Sleep(time.Microsecond)
			fmt.Println("func ", callSign, " writing")
		}
		mutex.Unlock()
	}
	go fun(1)
	go fun(2)
	time.Sleep(time.Millisecond)
}
