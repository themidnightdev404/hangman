package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func loadWords() []string {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	var words []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			words = append(words, strings.ToLower(line))
		}
	}

	if scanner.Err() != nil {
		fmt.Println("Ошибка чтения файла", scanner.Err())
	}

	return words
}

func getRandomWord(words []string) string {
	index := rand.Intn(len(words))
	return words[index]
}
