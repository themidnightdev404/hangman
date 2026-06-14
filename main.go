package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	words := loadWords()
	reader := bufio.NewReader(os.Stdin)
	for {
		ClearTerminal()

		fmt.Println("1 - Новая игра")
		fmt.Println("2 - Выйти")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "1" {
			fmt.Println("Начинаем игру")
			random := getRandomWord(words)
			PlayGame(random, reader)

			fmt.Println("\nНажмите Enter, чтобы вернуться в меню...")
			reader.ReadBytes('\n')
		} else if choice == "2" {
			fmt.Println("Выходим")
			os.Exit(0)
		} else {
			fmt.Println("⚠️ Упс! Такого пункта в меню нет. Попробуйте еще раз.")
			fmt.Println("\nНажмите Enter, чтобы продолжить...")
			reader.ReadBytes('\n')
		}
	}
}
