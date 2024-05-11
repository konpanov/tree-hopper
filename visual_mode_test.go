package main

import (
	"testing"
)

func TestVisualMovment(t *testing.T) {
	win := createWindowFromString("qwe\r\nasd", 80, 200)
	cursorRight(win)
	enterVisualMode(win)
	cursorDown(win)

	vo := win.visualOrigin

	if vo.char != 1 {
		temp := "Incorrect visual origin char %d != %d"
		t.Errorf(temp, vo.char, 1)
	}
	if vo.line != 0 {
		temp := "Incorrect visual origin line %d != %d"
		t.Errorf(temp, vo.char, 0)
	}
	if win.cursor.char != 1 {
		temp := "Incorrect cursor char %d != %d"
		t.Errorf(temp, win.cursor.char, 0)
	}
	if win.cursor.line != 1 {
		temp := "Incorrect cursor line %d != %d"
		t.Errorf(temp, win.cursor.line, 1)
	}
}
