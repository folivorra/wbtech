package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы, присутствующие и в первом, и во втором.
//Пример:
//A = {1,2,3}
//B = {2,3,4}
//Пересечение = {2,3}

func main() {
	a, b := []int{1, 2, 3}, []int{2, 3, 4}

	set := make(map[int]struct{})

	for _, v := range a {
		set[v] = struct{}{}
	}

	var res []int
	for _, v := range b {
		if _, exist := set[v]; exist {
			res = append(res, v)
		}
	}

	fmt.Println(res)
}
