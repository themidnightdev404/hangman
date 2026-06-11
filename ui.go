package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"unicode"
)

var STAGES = [7]string{
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

func PlayGame(word string) {
	game := NewGame(word)
	var input string
	for {
		ClearTerminal()

		fmt.Println(STAGES[game.Mistakes])

		printWordState(game.Word, game.Guessed)
		fmt.Println()

		fmt.Print("Использованные буквы: ")
		for letter := range game.UsedLetter {
			fmt.Print(string(letter) + " ")
		}
		fmt.Println()

		fmt.Print("Введите букву: ")
		fmt.Scan(&input)
		input = strings.ToLower(input)
		if len([]rune(input)) != 1 {
			fmt.Println("Пожалуйста, введите только одну букву")
			fmt.Println("Нажмите Enter, чтобы продолжить")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			continue
		}
		char := []rune(input)[0]
		if !unicode.Is(unicode.Cyrillic, char) {
			fmt.Println("Введите букву русского алфавита!")
			fmt.Println("Нажмите Enter, чтобы продолжить")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			continue
		}
		if game.UsedLetter[char] {
			fmt.Println("Вы уже вводили эту букву! Попробуйте другую.")
			fmt.Println("Нажмите Enter, чтобы продолжить")
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			bufio.NewReader(os.Stdin).ReadBytes('\n')
			continue
		}
		game.UsedLetter[char] = true
		game.Guess(input)
		if game.IsWin() {
			fmt.Println("Поздравляем! Слово \"" + game.Word + "\" угадано!")
			break
		}
		if game.IsLose() {
			ClearTerminal()
			fmt.Println(STAGES[game.Mistakes])
			printWordState(game.Word, game.Guessed)
			fmt.Println()
			fmt.Println("Не повезло! Правильный ответ:", game.Word)
			break
		}

	}

}

func printWordState(word string, guessed []string) {
	for _, letter := range word {
		guessedStr := strings.Join(guessed, "")
		if strings.Contains(guessedStr, string(letter)) {
			fmt.Print(string(letter) + " ")
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func ClearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
