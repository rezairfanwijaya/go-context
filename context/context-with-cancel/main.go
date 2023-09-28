package main

import (
	"context"
	"log"
	"runtime"
	"sync"
	"time"
)

func createCounter(cancel context.Context, group *sync.WaitGroup, mutex *sync.Mutex) chan int {
	group.Add(1)

	finalCounter := make(chan int)
	counter := 1
	go func() {
		defer close(finalCounter)
		for {
			select {
			case <-cancel.Done():
				return // membatalkan goroutine
			default:
				finalCounter <- counter
				counter++
			}
		}
	}()
	group.Done()

	return finalCounter
}

func main() {
	log.Printf("start %v goroutine", runtime.NumGoroutine())

	group := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	background := context.Background()
	contextCancel, cancel := context.WithCancel(background)

	res := createCounter(contextCancel, group, mutex)
	log.Printf("proses %v goroutine", runtime.NumGoroutine())

	for value := range res {
		log.Println("val : ", value)
		if value == 5 {
			break
		}
	}
	cancel() // proses pengiriman sinyal cancel
	group.Wait()

	<-time.After(1 * time.Minute)
	log.Printf("remain %v goroutine", runtime.NumGoroutine())
}
