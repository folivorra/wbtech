package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep, которая приостанавливает выполнение текущей горутины.

func main() {
	now := time.Now()
	sleep(2 * time.Second)
	fmt.Println(time.Since(now))
}

func sleep(duration time.Duration) {
	done := make(chan struct{})

	go func() {
		<-time.After(duration)
		done <- struct{}{}
	}()

	<-done

	//<-time.After(duration) - пример как бы тоже работало
}

// только я не понял зачем указан способ с горутиной и каналом, если можно просто указать <- time.After(duration) и он будет также работать
