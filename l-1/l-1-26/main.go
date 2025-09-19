package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).
// Вывод: true, если все символы уникальны, false, если есть повторения. Проверка должна быть регистронезависимой,
// т.е. символы в разных регистрах считать одинаковыми.

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(checkUnique(str))
}

func checkUnique(str string) bool {
	uniqueLetter := make(map[rune]struct{})

	for _, v := range strings.ToLower(str) {
		if _, ok := uniqueLetter[v]; ok {
			return false
		}
		uniqueLetter[v] = struct{}{}
	}

	return true
}
