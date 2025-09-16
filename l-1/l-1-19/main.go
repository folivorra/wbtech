package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Разработать программу, которая переворачивает подаваемую на вход строку.
//Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

func main() {
	var str string

	reader := bufio.NewReader(os.Stdin)
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	
	runes := []rune(str)
	for i := len(runes) - 1; i >= 0; i-- {
		fmt.Print(string(runes[i]))
	}
}
