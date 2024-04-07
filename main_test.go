package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func assertContent(t *testing.T, s tcell.Screen, x, y int, expected rune) {
	r, _, _, _ := s.GetContent(x, y)
	if r != expected {
		t.Fail()
	}
}

func TestDrawWindow(t *testing.T) {
	win := createWindow("assets/hello_world.txt")
	drawWindow(win)
	assertContent(t, win.screen, 0, 0, 'H')
	assertContent(t, win.screen, 1, 0, 'e')
	assertContent(t, win.screen, 2, 0, 'l')
	assertContent(t, win.screen, 3, 0, 'l')
	assertContent(t, win.screen, 4, 0, 'o')
	assertContent(t, win.screen, 5, 0, ',')
	assertContent(t, win.screen, 6, 0, ' ')
	assertContent(t, win.screen, 7, 0, 'W')
	assertContent(t, win.screen, 8, 0, 'o')
	assertContent(t, win.screen, 9, 0, 'r')
	assertContent(t, win.screen, 10, 0, 'l')
	assertContent(t, win.screen, 11, 0, 'd')
	assertContent(t, win.screen, 12, 0, '!')
	assertContent(t, win.screen, 13, 0, ' ')
	assertContent(t, win.screen, 14, 0, ' ')
}
