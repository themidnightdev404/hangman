package main

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
)

var stages = [7]string{
	`
  +---+

  |   |
      |
      |
      |
      |
=========`,
	`
  +---+

  |   |
  O   |
      |
      |
      |
=========`,
	`
  +---+

  |   |
  O   |

  |   |
  |   |
      |
=========`,
	`
  +---+

  |   |
  O   |
 /|   |

  |   |
      |
=========`,
	`
  +---+

  |   |
  O   |
 /|\  |

  |   |
      |
=========`,
	`
  +---+

  |   |
  O   |
 /|\  |

  |   |
 /    |
=========`,
	`
  +---+

  |   |
  O   |
 /|\  |

  |   |
 / \  |
=========`,
}

func PlayGame(word string, reader *bufio.Reader) {
	game := NewGame(word)
	var input string
	for {
		ClearTerminal()

		fmt.Println(stages[game.Mistakes])

		printWordState(game.Word, game.UsedLetter)
		fmt.Println()

		fmt.Print("Использованные буквы: ")
		for letter := range game.UsedLetter {
			fmt.Print(string(letter) + " ")
		}
		fmt.Println()

		fmt.Print("Введите букву: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ToLower(input)
		if len([]rune(input)) != 1 {
			fmt.Println("Пожалуйста, введите только одну букву")
			fmt.Println("Нажмите Enter, чтобы продолжить...")
			reader.ReadBytes('\n')
			continue
		}
		char := []rune(input)[0]
		if !unicode.Is(unicode.Cyrillic, char) {
			fmt.Println("Введите букву русского алфавита!")
			fmt.Println("Нажмите Enter, чтобы продолжить...")
			reader.ReadBytes('\n')
			continue
		}
		if game.UsedLetter[char] {
			fmt.Println("Вы уже вводили эту букву! Попробуйте другую.")
			fmt.Println("Нажмите Enter, чтобы продолжить...")
			reader.ReadBytes('\n')
			continue
		}
		game.Guess(char)
		if game.IsWin() {
			fmt.Println("Поздравляем! Слово \"" + game.Word + "\" угадано!")
			break
		}
		if game.IsLose() {
			ClearTerminal()
			fmt.Println(stages[game.Mistakes])
			printWordState(game.Word, game.UsedLetter)
			fmt.Println()
			fmt.Println("Не повезло! Правильный ответ:", game.Word)
			break
		}

	}

}

func printWordState(word string, guessed map[rune]bool) {
	for _, letter := range word {
		if guessed[letter] {
			fmt.Print(string(letter) + " ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}
