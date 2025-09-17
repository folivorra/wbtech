package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

func main() {
	var a, b string
	fmt.Scan(&a)
	fmt.Scan(&b)

	aInt := big.NewInt(0)
	bInt := big.NewInt(0)

	aInt.SetString(a, 10)
	bInt.SetString(b, 10)

	tmp := big.NewInt(0)

	fmt.Println("a+b:", tmp.Add(aInt, bInt).String(), "\na-b:", tmp.Sub(aInt, bInt).String(), "\na*b:", tmp.Mul(aInt, bInt).String(), "\na/b:", tmp.Div(aInt, bInt).String())
}
