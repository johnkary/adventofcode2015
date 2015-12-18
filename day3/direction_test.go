package main

import (
	"testing"
)

func TestFromStringFiltersInvalidCharacters(t *testing.T) {
	cases := []struct {
		input  string
		length int
	}{
		{"", 0},
		{"asdf", 0},
		{" ^v<> g", 4},
		{" ^ ^< > vv^ &\"", 7},
	}
	for _, c := range cases {
		moves := FromString(c.input)
		if len(moves) != c.length {
			t.Errorf("Expected %d but got %d with input: %q", c.length, len(moves), c.input)
		}
	}
}

func TestMovement(t *testing.T) {
	cases := []struct{
		input []Mover
		expected Coord
	}{
		{
			[]Mover{Right(), Right()},
			Coord{0,2}
		}
	}
	for _, c := range cases {

		if coordinate == c.expected {
			t.Errorf("Expected %v but got %v with input: %s", c.expected, coordinate, c.input)
		}
	}
}
