package main

import "fmt"

//Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var i int
	fmt.Scan(&i)

	fmt.Println("элементы массива:", arr)
	fmt.Println("длина массива:", len(arr))

	arr = deleteElem(arr, i)
	fmt.Println("элементы массива после удаления:", arr)
	fmt.Println("длина массива после удаления:", len(arr))
}

func deleteElem(arr []int, i int) []int {
	if i < 0 || i >= len(arr) {
		return []int{}
	}

	copy(arr[i:], arr[i+1:])

	return arr[:len(arr)-1]
}
