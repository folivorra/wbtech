package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Реализовать постоянную запись данных в канал (в главной горутине).
//Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
//Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.

func main() {
	var workersNum int
	fmt.Scan(&workersNum)

	ch := make(chan int)

	for i := 0; i < workersNum; i++ {
		go worker(i+1, ch)
	}

	for {
		ch <- rand.Intn(100)
		time.Sleep(200 * time.Millisecond)
	}
}

func worker(id int, in <-chan int) {
	for n := range in {
		fmt.Printf("worker %d: %d\n", id, n)
	}
}
