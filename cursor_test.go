package main

import (
	"testing"
)

func TestCursorRight(t *testing.T) {
	win := createWindowFromString("qwe\r\nasd", 80, 200)
	cursorRight(win)
	expected := CursorPos{0, 1}
	if *win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorRightDoesNotSkipLines(t *testing.T) {
	win := createWindowFromString("qwe\r\nasd", 80, 200)
	cursorRight(win)
	cursorRight(win)
	cursorRight(win)
	expected := CursorPos{0, 2}
	if *win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorLeftDoesNotSkipLines(t *testing.T) {
	win := createWindowFromString("qwe\r\nasd", 80, 200)
	cursorDown(win)
	cursorRight(win)
	cursorLeft(win)
	cursorLeft(win)
	cursorLeft(win)
	cursorLeft(win)
	expected := CursorPos{1, 0}
	if *win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorDown(t *testing.T) {
	win := createWindowFromString("qwe\r\nasd", 80, 200)
	cursorDown(win)
	expected := CursorPos{1, 0}
	if *win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorDownAndUp(t *testing.T) {
	win := createWindowFromString("qwe\r\nasd", 80, 200)
	cursorDown(win)
	cursorUp(win)
	expected := CursorPos{0, 0}
	if *win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorDownUpDown(t *testing.T) {
	win := createWindowFromString("qwe\r\nasd", 80, 200)
	cursorDown(win)
	cursorUp(win)
	cursorDown(win)
	expected := CursorPos{1, 0}
	if *win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}
