package main

import (
    "strings"
)

func part2(input string) int {
    area := 0
    for _, dimensions := range strings.Split(input, "\n") {
        area += getRibbonLength(dimensions)
    }

    return area
}

func getRibbonLength(dimensions string) int {
    p := Present{boxFromCoord(dimensions)}

    return p.RibbonWrapLength() + p.RibbonBowLength()
}
