package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значения квадратов чисел, взятых из массива [2,4,6,8,10], и выведет результаты в stdout.

func main() {
	mas := []int{
		2, 4, 6, 8, 10,
	}

	wg := &sync.WaitGroup{}

	for _, v := range mas {
		wg.Add(1)
		go square(v, wg)
	}

	wg.Wait()
}

func square(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(n, ":", n*n)
}
