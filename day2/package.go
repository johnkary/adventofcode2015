package main

import (
    "sort"
)

type Present struct {
    Box
}

func (p *Present) WrappingPaperArea() int {
    return p.Area() + p.padding()
}

func (p *Present) padding() int {
    sides := make([]int, 3)
    sides[0] = p.l * p.w
    sides[1] = p.w * p.h
    sides[2] = p.h * p.l

    sort.Ints(sides)

    return sides[0]
}
