package main

import (
	"testing"
)

func TestHasVisited(t *testing.T) {
	cases := []struct {
		visited    []Coord
		search     Coord
		hasVisited bool
	}{
		{nil, Coord{0, 0}, true},                               // Home always visited
		{nil, Coord{0, 1}, false},                              // Unvisited coord
		{[]Coord{Coord{0, 1}}, Coord{0, 1}, true},              // First move
		{[]Coord{Coord{0, 1}, Coord{1, 1}}, Coord{1, 1}, true}, // Second move
	}
	for _, c := range cases {
		houses := NewVisitedHouses(c.visited)

		visited := houses.HasVisited(c.search)
		if visited != c.hasVisited {
			t.Errorf("Expected %t but got %t with input: %#v", c.hasVisited, visited, houses)
		}
	}
}

func TestCount(t *testing.T) {
	cases := []struct {
		visited []Coord
		count   int
	}{
		{nil, 1},                                            // Home is always visited
		{[]Coord{Coord{0, 1}}, 2},                           // One move
		{[]Coord{Coord{0, 1}, Coord{1, 1}}, 3},              // Two moves
		{[]Coord{Coord{0, 1}, Coord{1, 1}, Coord{0, 1}}, 3}, // Multiple of same moves
	}
	for _, c := range cases {
		houses := NewVisitedHouses(c.visited)
		count := houses.Count()
		if count != c.count {
			t.Errorf("Expected %d but got %d with input: %#v", c.count, count, houses)
		}
	}
}
