package main

import (
    "strings"
)

func part1(input string) int {
    area := 0
    for _, dimensions := range strings.Split(input, "\n") {
        area += getPaperArea(dimensions)
    }

    return area
}

func getPaperArea(dimensions string) int {
    present := Present{boxFromCoord(dimensions)}

    return present.WrappingPaperArea()
}
