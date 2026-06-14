package main

import "strings"

const maxMistakes = 6

type Game struct {
	Word       string
	Mistakes   int
	UsedLetter map[rune]bool
}

func NewGame(
	word string,
) Game {
	return Game{
		Word:       strings.ToLower(word),
		UsedLetter: map[rune]bool{},
	}
}

func (g *Game) Guess(char rune) bool {
	g.UsedLetter[char] = true
	contains := strings.ContainsRune(g.Word, char)
	if !contains {
		g.Mistakes++
	}
	return contains
}

func (g *Game) IsWin() bool {
	for _, letter := range g.Word {
		if !g.UsedLetter[letter] {
			return false
		}
	}
	return true
}

func (g *Game) IsLose() bool {
	return g.Mistakes >= maxMistakes
}
