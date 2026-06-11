package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := loadWords()

	for {
		ClearTerminal()
		var game string

		fmt.Println("1 - Новая игра")
		fmt.Println("2 - Выйти")
		fmt.Scan(&game)

		if game == "1" {
			fmt.Println("Начинаем игру")
			random := getRandomWord(words)
			PlayGame(random)

			fmt.Println("\nНажмите Enter, чтобы вернуться в меню...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			bufio.NewReader(os.Stdin).ReadBytes('\n')
		} else {
			fmt.Println("Выходим")
			os.Exit(0)
		}
	}
}
