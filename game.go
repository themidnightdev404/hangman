package main

import "strings"

const maxMistakes = 6

type Game struct {
	Word       string
	Mistakes   int
	Guessed    []string
	UsedLetter map[rune]bool
}

func NewGame(
	word string,
) Game {
	return Game{
		Word:       word,
		UsedLetter: map[rune]bool{},
	}
}

func (g *Game) Guess(newGuess string) bool {
	contains := strings.Contains(g.Word, newGuess)
	if contains {
		g.Guessed = append(g.Guessed, newGuess)
	} else {
		g.Mistakes++
	}
	return contains
}

func (g *Game) IsWin() bool {
	for _, letter := range g.Word {
		guessedStr := strings.Join(g.Guessed, "")
		if !strings.Contains(guessedStr, string(letter)) {
			return false
		}
	}
	return true
}

func (g *Game) IsLose() bool {
	return g.Mistakes == maxMistakes
}
