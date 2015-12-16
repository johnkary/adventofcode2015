package main

import (
    "testing"
)

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

func TestRibbonWrapLength(t *testing.T) {
    cases := []struct {
        present Present
        want int
    }{
        {Present{Box{l:2, w:3, h:4}}, 10},
        {Present{Box{l:1, w:1, h:10}}, 4},
        {Present{Box{l:3, w:2, h:1}}, 6}, // I think
    }
    for _, c := range cases {
        length := c.present.RibbonWrapLength()
        if length != c.want {
            t.Errorf("Expected %d but got %d with input: %#v", c.want, length, c.present)
        }
    }
}

func TestRibbonBowLength(t *testing.T) {
    cases := []struct {
        present Present
        want int
    }{
        {Present{Box{l:2, w:3, h:4}}, 24},
        {Present{Box{l:1, w:1, h:10}}, 10},
        {Present{Box{l:3, w:2, h:1}}, 6},
    }
    for _, c := range cases {
        length := c.present.RibbonBowLength()
        if length != c.want {
            t.Errorf("Expected %d but got %d with input: %#v", c.want, length, c.present)
        }
    }
}
