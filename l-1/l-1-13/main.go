package main

import "fmt"

//Поменять местами два числа без использования временной переменной.
//Подсказка: примените сложение/вычитание или XOR-обмен.

func main() {
	addAndSub(123, 345)

	xor(65, 12)
}

func addAndSub(a, b int) {
	fmt.Println("-----сложение и вычитание-----")
	fmt.Printf("исходные a = %d, b = %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("поменяные a = %d, b = %d\n\n", a, b)
}

func xor(a, b int) {
	fmt.Println("----------XOR-обмен----------")
	fmt.Printf("исходные a = %d, b = %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("поменяные a = %d, b = %d\n\n", a, b)
}
