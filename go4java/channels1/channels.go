package main

import (
	"log"
	"time"
)

func main() {
	channel := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		close(channel)
	}()
	go func() {
		for i := range channel { // read indefinitely
			log.Println("receiver:", i)
		}
	}()
	time.Sleep(time.Millisecond)
}
