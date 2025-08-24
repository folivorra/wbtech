package main

import (
	"fmt"
	"sync"
)

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
