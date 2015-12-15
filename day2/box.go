package main

import (
    "strings"
    "strconv"
)

type Box struct {
    l int
    w int
    h int
}

func boxFromCoord(input string) Box {
    fields := strings.Split(input, "x")
    length, err := strconv.Atoi(fields[0])
    check(err)
    width, err := strconv.Atoi(fields[1])
    check(err)
    height, err := strconv.Atoi(fields[2])
    check(err)

    return Box{l:length, w:width, h:height}
}

func (b *Box) Area() int {
    return 2*b.l*b.w + 2*b.w*b.h + 2*b.h*b.l
}
