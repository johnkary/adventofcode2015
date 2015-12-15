package main

import "testing"

func TestWrappingPaperArea(t *testing.T) {
    cases := []struct {
        present Present
        want int
    }{
        {Present{Box{l:2, w:3, h:4}}, 58},
        {Present{Box{l:1, w:1, h:10}}, 43},
        {Present{Box{l:3, w:2, h:1}}, 24},
    }
    for _, c := range cases {
        area := c.present.WrappingPaperArea()
        if area != c.want {
            t.Errorf("Expected %d but got %d with input: %#v", c.want, area, c.present)
        }
    }
}

// func TestPart1(t *testing.T) {
//     cases := []struct {
//         in string
//         want int
//     }{
//         {"2x3x4", 58},
//         {"1x1x10", 43},
//     }
//     for _, c := range cases {
//         got := part1(c.in)
//         if got != c.want {
//             t.Errorf("Expected %q but got %q with input: %q", c.want, got, c.in)
//         }
//     }
// }
