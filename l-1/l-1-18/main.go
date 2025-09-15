package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
//По завершению программы структура должна выводить итоговое значение счётчика.

type Counter struct {
	value atomic.Int64
}

func main() {
	counter := Counter{}

	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(counter *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			counter.value.Add(1)
		}(&counter, &wg)
	}

	wg.Wait()
	fmt.Println(counter.value.Load())
}
