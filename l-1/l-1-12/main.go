package main

import "fmt"

//Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.
//Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func main() {
	set := make(map[string]struct{})
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	var setSlice []string
	for _, word := range words {
		if _, exist := set[word]; !exist {
			setSlice = append(setSlice, word)
		}
		set[word] = struct{}{}
	}

	fmt.Println(setSlice)
}
