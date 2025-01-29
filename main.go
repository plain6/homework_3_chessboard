package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var board strings.Builder

	fmt.Print("Программа \"Шахматная доска\".\nВведите размер доски: ")
	fmt.Scanln(&n)

	if n > 0 {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if (i+j)%2 == 0 {
					board.WriteString("# ")
				} else {
					board.WriteString("  ")
				}
			}

			board.WriteString("\n")
		}

		fmt.Println(board.String())
	} else {
		fmt.Println("Введён неправильный размер доски.")
	}
}
