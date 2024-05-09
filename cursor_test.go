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

func TestSlideWindowOnDown(t *testing.T) {
	win := createWindowFromString(createNLines(100), 80, 40)
	assertIntEqual(t, win.topLine, 0)
	for i := 0; i < 40; i++ {
		cursorDown(win)
	}
	assertIntEqual(t, win.topLine, 1)
}

func TestSlideWindowUp(t *testing.T) {
	win := createWindowFromString(createNLines(100), 80, 40)
	for i := 0; i < 60; i++ {
		cursorDown(win)
	}
	assertIntEqual(t, win.topLine, 21)
	for i := 0; i < 60; i++ {
		cursorDown(win)
	}
}

func TestCursorStopAtLastLine(t *testing.T) {
	win := createWindow("assets/tests/42lines.txt", 80, 40)
	for i := 0; i < 150; i++ {
		cursorDown(win)
	}
	assertIntEqual(t, win.cursor.line, 41)
}
