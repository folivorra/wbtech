package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

//Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
//Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.

func main() {
	var workersNum int
	fmt.Scan(&workersNum)

	// использую сигнальный канал для отлова sigint
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)

	// создаю контекст с отменой для управления продолжительностью жизни горутин, так как этот способ
	// используется в большинстве библиотек и с ним удобнее работать, чем с каналом для оповещения
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// waitgroup для того чтобы дождаться завершения всех горутин-воркеров и генератора данных
	wg := &sync.WaitGroup{}

	ch := make(chan int)

	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go worker(ctx, wg, i+1, ch)
	}

	wg.Add(1)
	go generateData(ctx, wg, ch)

	<-shutdown
	cancel()

	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup, id int, in <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d done\n", id)
			return
		case n := <-in:
			fmt.Printf("worker %d: %d\n", id, n)
		}
	}
}

func generateData(ctx context.Context, wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	defer close(ch)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("generator done")
			return
		case ch <- rand.Intn(100):
			time.Sleep(200 * time.Millisecond)
		}
	}
}
