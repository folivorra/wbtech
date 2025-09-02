package main

import "fmt"

//Реализовать алгоритм бинарного поиска встроенными методами языка.
//Функция должна принимать отсортированный слайс и искомый элемент,
//возвращать индекс элемента или -1, если элемент не найден.

func main() {
	arr := []int{1, 3, 8, 23, 66, 3333, 21234}
	fmt.Println(binarySearch(arr, 3333))
	fmt.Println(binarySearch(arr, 1))
	fmt.Println(binarySearch(arr, 0))

}

func binarySearch(arr []int, x int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if x == arr[mid] {
			return mid
		} else if x > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
