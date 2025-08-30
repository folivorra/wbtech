package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.
//Классические подходы: выход по условию, через канал уведомления,
//через контекст, прекращение работы runtime.Goexit() и др.

func main() {
	// все примеры завершают себя самостоятельно без взаимодействия с
	// main горутиной для упрощения примеров

	wg := sync.WaitGroup{}
	wg.Add(5)

	// 1
	go func() {
		defer wg.Done()
		condition()
	}()

	//2
	go func() {
		defer wg.Done()
		channel()
	}()

	//3
	go func() {
		defer wg.Done()
		ctx()
	}()

	//4
	go func() {
		defer wg.Done()
		exit()
	}()

	//5
	go func() {
		defer wg.Done()
		closeChan()
	}()

	wg.Wait()
}

func condition() {
	f := false

	for {
		if f {
			fmt.Println("1 close")
			return
		}
		f = true
	}
}

func channel() {
	done := make(chan struct{}, 1)

	for {
		select {
		case <-done:
			fmt.Println("2 close")
			return
		default:
		}
		done <- struct{}{}
	}
}

func ctx() {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		select {
		case <-c.Done():
			fmt.Println("3 close")
			return
		default:
		}
		cancel()
	}
}

func exit() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("4 close")
		runtime.Goexit()
	}
}

func closeChan() {
	val := make(chan struct{}, 1)
	val <- struct{}{}

	for {
		_, ok := <-val
		if !ok {
			fmt.Println("5 close")
			return
		}

		close(val)
	}
}
