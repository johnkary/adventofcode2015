package main

import (
    "sort"
)

type Present struct {
    Box
}

func (p *Present) WrappingPaperArea() int {
    return p.Area() + p.paperPadding()
}

func (p *Present) paperPadding() int {
    sides := make([]int, 3)
    sides[0] = p.l * p.w
    sides[1] = p.w * p.h
    sides[2] = p.h * p.l

    sort.Ints(sides)

    return sides[0]
}

func (p *Present) RibbonWrapLength() int {
    sides := make([]int, 3)
    sides[0] = p.l
    sides[1] = p.w
    sides[2] = p.h

    sort.Ints(sides)

    return sides[0] + sides[0] + sides[1] + sides[1]
}

func (p *Present) RibbonBowLength() int {
    return p.l * p.w * p.h
}
