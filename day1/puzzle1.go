package puzzle1

import (
    "strings"
)

func Puzzle1(instructions string) int {
    up := strings.Count(instructions, "(")
    down := strings.Count(instructions, ")")

    return up - down
}
