package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

// Переворот текста
func reverseHello(text string) string {
	return reverse.String(text)
}

// Запуск программы
func main() {
	// Вовод в консоль результат
	fmt.Println(reverseHello("Hello, DIASOFT!"))
}
