package main

import "fmt"

//Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.
//Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).

func main() {
	var (
		a, i, bit int64
	)

	fmt.Scan(&a, &i, &bit)
	if bit == 1 {
		a = a | (1 << i)
	} else if bit == 0 {
		a = a & ^(1 << i)
	}

	fmt.Println(a)
}
