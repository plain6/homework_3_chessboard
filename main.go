package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var board strings.Builder
	var toggle bool

	fmt.Print("Программа \"Шахматная доска\".\nВведите размер доски: ")
	fmt.Scanln(&n)

	if n > 0 {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if toggle {
					board.WriteString("  ")
					toggle = false
				} else {
					board.WriteString("# ")
					toggle = true
				}
			}

			board.WriteString("\n")
			if n%2 == 0 {
				if toggle {
					toggle = false
				} else {
					toggle = true
				}
			}
		}

		fmt.Println(board.String())
	} else {
		fmt.Println("Введён неправильный размер доски.")
	}
}
