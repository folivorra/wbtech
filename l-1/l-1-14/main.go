package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}).
//Типы, которые нужно распознавать: int, string, bool, chan (канал).
//Подсказка: оператор типа switch v.(type) поможет в решении.

func main() {
	mas := []interface{}{1, "hello", true, make(chan map[int]string), complex(1.0, -2.0)}

	for _, m := range mas {
		resolver(m)
	}
}

func resolver(variable interface{}) {
	switch variable.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		t := reflect.TypeOf(variable)
		if t.Kind() == reflect.Chan {
			fmt.Println("chan")
		} else {
			fmt.Println("unknown")
		}
	}
}
