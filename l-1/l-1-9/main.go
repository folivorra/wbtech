package main

import (
	"fmt"
	"sync"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.
//После этого данные из второго канала должны выводиться в stdout. То есть, организуйте конвейер из двух этапов с горутинами:
//генерация чисел и их обработка. Убедитесь, что чтение из второго канала корректно завершается.

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		input(ch1)
	}()

	go func() {
		defer wg.Done()
		square(ch1, ch2)
	}()

	go func() {
		defer wg.Done()
		stdout(ch2)
	}()

	wg.Wait()
}

func input(out chan<- int) {
	mas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for _, v := range mas {
		out <- v
	}

	close(out)
}

func square(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * 2
	}

	close(out)
}

func stdout(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
