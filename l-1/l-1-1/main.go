package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	name string
	age  int
	sex  string
}

func (h *Human) Name() string {
	return h.name
}

func (h *Human) Age() int {
	return h.age
}

func (h *Human) Sex() string {
	return h.sex
}

func (h *Human) SetAge(age int) {
	h.age = age
}

func (h *Human) SetSex(s string) {
	h.sex = s
}

type Action struct {
	Human
}

func main() {
	a := Action{
		Human: Human{name: "Sasha", age: 18, sex: "male"},
	}

	fmt.Println("name:", a.Name())
	fmt.Println("age:", a.Age())

	a.SetAge(30)
	a.SetSex("female")

	fmt.Println("updated age:", a.Age())
	fmt.Println("updated sex:", a.Sex())
}
