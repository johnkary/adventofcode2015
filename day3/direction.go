package main

import (
	"strings"
	"unicode/utf8"
)

type Mover interface {
	Move() (int, int)
}

type X struct {
	mod int
}

func (m X) Move(from Coord) (int, int) {
	return m.mod, 0
}

type Y struct {
	mod int
}

func (m Y) Move() (int, int) {
	return 0, m.mod
}

func Left() X {
	return X{-1}
}
func Right() X {
	return X{1}
}
func Up() Y {
	return Y{1}
}
func Down() Y {
	return Y{-1}
}

func FromString(input string) []Mover {
	// TODO: try doing... map[string]Mover{"^":Up(), ">":Right(), "v":Down(), "<":Left()}
	lookup := make(map[string]Mover)
	lookup["^"] = Up()
	lookup[">"] = Right()
	lookup["v"] = Down()
	lookup["<"] = Left()

	// Translate to moves
	translate := make(map[rune]Mover)
	for direction, mover := range lookup {
		r, _ := utf8.DecodeRuneInString(direction)
		translate[r] = mover
	}

	// Remove invalid input
	validRunes := keys(translate)
	validInput := strings.Map(func(s rune) rune {
		for _, valid := range validRunes {
			if s == valid {
				return s
			}
		}
		return -1
	}, input)

	// Translate valid to moves
	moves := make([]Mover, len(validInput))
	for i, r := range validInput {
		moves[i] = translate[r]
	}

	return moves
}

// Return keys of the given map
func keys(m map[rune]Mover) (keys []rune) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
