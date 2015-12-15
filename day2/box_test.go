package main

import "testing"

func TestBoxFromCoord(t *testing.T) {
    cases := []struct {
        in string
        l int
        w int
        h int
    }{
        {"2x3x4", 2, 3, 4},
        {"1x1x10", 1, 1, 10},
    }
    for _, c := range cases {
        box := boxFromCoord(c.in)
        if box.l != c.l {
            t.Errorf("Length: Expected %d but got %d with input: %q", c.l, box.l, c.in)
        }
        if box.w != c.w {
            t.Errorf("Width: Expected %d but got %d with input: %q", c.w, box.w, c.in)
        }
        if box.h != c.h {
            t.Errorf("Height: Expected %d but got %d with input: %q", c.h, box.h, c.in)
        }
    }
}

func TestArea(t *testing.T) {
    cases := []struct {
        box Box
        want int
    }{
        {Box{l:2, w:3, h:4}, 52},
        {Box{l:1, w:1, h:10}, 42},
    }
    for _, c := range cases {
        if c.box.Area() != c.want {
            t.Errorf("Expected %d but got %d with input: %#v", c.want, c.box.Area(), c.box)
        }
    }
}
