package main

import (
	"log"
	"sync"
)

func main() {
	channel, wg := make(chan int, 10), sync.WaitGroup{}
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		close(channel)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range channel {
			log.Println("receiver:", i)
		}
	}()
	wg.Wait()
}
