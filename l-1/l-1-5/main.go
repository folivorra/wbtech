package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала – читать эти значения.
//По истечении N секунд программа должна завершаться.

func main() {
	var timeoutSeconds int
	fmt.Scan(&timeoutSeconds)

	ch := make(chan int)

	wg := &sync.WaitGroup{}

	timeout, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	wg.Add(2)
	go generateData(timeout, wg, ch)
	go printData(timeout, wg, ch)

	wg.Wait()

	fmt.Printf("timeout: %ds\n", timeoutSeconds)
}

func generateData(ctx context.Context, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	defer close(ch)
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- rand.Intn(100):
		}
	}
}

func printData(ctx context.Context, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-ch:
			fmt.Println(data)
		}
	}
}
