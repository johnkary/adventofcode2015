package main

type VisitedHouses struct {
	visited []Coord
}

func NewVisitedHouses(visited []Coord) *VisitedHouses {
	h := new(VisitedHouses)
	h.Visit(Coord{x: 0, y: 0})

	for _, coord := range visited {
		h.Visit(coord)
	}

	return h
}

func (v *VisitedHouses) Visit(c Coord) {
	if false == v.HasVisited(c) {
		v.visited = append(v.visited, c)
	}
}

func (v *VisitedHouses) HasVisited(c Coord) bool {
	for _, visit := range v.visited {
		if c == visit {
			return true
		}
	}
	return false
}

func (v *VisitedHouses) Count() int {
	return len(v.visited)
}
