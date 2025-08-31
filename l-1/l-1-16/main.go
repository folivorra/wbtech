package main

import "fmt"

//Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.
//Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел. Для выбора опорного элемента можно взять середину или первый элемент.

func main() {
	arr := []int{4, 7, 2, 1, 7, 4, 2, 5, 3, 1, 4, 7, 8, 4, 3, 10, 1, 2312, 432}

	fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
	i, j := 0, len(arr)-1
	pivot := findPivot(arr)
	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	if j > 0 {
		quickSort(arr[:j+1])
	}
	if i < len(arr)-1 {
		quickSort(arr[i:])
	}
	return arr
}

func findPivot(nums []int) int {
	mid := len(nums) / 2
	a, b, c := nums[0], nums[mid], nums[len(nums)-1]
	if (a > b) != (a > c) {
		return a
	}
	if (b > a) != (b > c) {
		return b
	}
	return c
}
