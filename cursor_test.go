package main

import (
	"testing"
)

func TestCursorRight(t *testing.T) {
	win := createWindow("assets/hello_world.txt")
	win.content = "qwe\nasd"
	cursorRight(win)
	expected := 1
	if win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorRightDoesNotSkipLines(t *testing.T) {
	win := createWindow("assets/hello_world.txt")
	win.content = "qwe\nasd"
	cursorRight(win)
	cursorRight(win)
	cursorRight(win)
	expected := 2
	if win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorLeftDoesNotSkipLines(t *testing.T) {
	win := createWindow("assets/hello_world.txt")
	win.content = "qwe\nasd"
	win.cursor = 7
	cursorLeft(win)
	cursorLeft(win)
	cursorLeft(win)
	cursorLeft(win)
	cursorLeft(win)
	cursorLeft(win)
	expected := 4
	if win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorDown(t *testing.T) {
	win := createWindow("assets/hello_world.txt")
	win.content = "qwe\nasd"
	cursorDown(win)
	expected := 4
	if win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorDownAndUp(t *testing.T) {
	win := createWindow("assets/hello_world.txt")
	win.content = "qwe\nasd"
	cursorDown(win)
	cursorUp(win)
	expected := 2
	if win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorDownUpDown(t *testing.T) {
	win := createWindow("assets/hello_world.txt")
	win.content = "qwe\nasd"
	cursorDown(win)
	cursorUp(win)
	cursorDown(win)
	expected := 4
	if win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorDownOnLF(t *testing.T) {
	win := createWindow("assets/hello_world.txt")
	win.content = "qwe\n\nasd"
	cursorDown(win)
	expected := 4
	if win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorUpOnDoubleLF(t *testing.T) {
	win := createWindow("assets/hello_world.txt")
	win.content = "qwe\n\nasd"
	cursorDown(win)
	cursorDown(win)
	cursorUp(win)
	expected := 4
	if win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorDownOnCRLF(t *testing.T) {
	win := createWindow("assets/hello_world.txt")
	win.content = "qwe\r\n\r\nasd"
	cursorDown(win)
	expected := 5
	if win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorUpOnDoubleCRLF(t *testing.T) {
	win := createWindow("assets/hello_world.txt")
	win.content = "qwe\r\n\r\nasd"
	cursorDown(win)
	cursorDown(win)
	cursorUp(win)
	expected := 5
	if win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}
