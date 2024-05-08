package main

import (
	"testing"
)

func TestInsertSymbol(t *testing.T) {
	win := createWindowFromString("", 80, 200)
	enterInsertMode(win)
	content := "h"
	insertUnderCursor(win, content)
	expected := content
	if win.lines[0] != expected {
		t.Errorf("%d != %s", win.cursor, expected)
	}
}

func TestCursorPosAfterInsert(t *testing.T) {
	win := createWindowFromString("", 80, 200)
	enterInsertMode(win)
	content := "h"
	insertUnderCursor(win, content)
	expected := CursorPos{0, 1}
	if *win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorPosAfterMultipleInserts(t *testing.T) {
	win := createWindowFromString("", 80, 200)
	enterInsertMode(win)
	content := "h"
	insertUnderCursor(win, content)
	insertUnderCursor(win, content)
	insertUnderCursor(win, content)
	expected := CursorPos{0, 3}
	if *win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestCursorPosInNormalModeAfterInsert(t *testing.T) {
	win := createWindowFromString("", 80, 200)
	enterInsertMode(win)
	content := "h"
	insertUnderCursor(win, content)
	enterNormalMode(win)
	expected := CursorPos{0, 0}
	if *win.cursor != expected {
		t.Errorf("%d != %d", win.cursor, expected)
	}
}

func TestRemoval(t *testing.T) {
	win := createWindowFromString("abcd", 80, 200)
	cursorRight(win)
	cursorRight(win)
	cursorRight(win)
	enterInsertMode(win)
	removeUnderCursor(win)
	enterNormalMode(win)
	expectedCursor := CursorPos{0, 1}
	if *win.cursor != expectedCursor {
		t.Errorf("%d != %d", win.cursor, expectedCursor)
	}
	expectedLine := "abd"
	if win.lines[0] != expectedLine {
		t.Errorf("%d != %s", win.cursor, expectedLine)
	}
}

func TestRemovalOnFirstSymbol(t *testing.T) {
	win := createWindowFromString("abcd", 80, 200)
	enterInsertMode(win)
	removeUnderCursor(win)
	enterNormalMode(win)
	expectedCursor := CursorPos{0, 0}
	if *win.cursor != expectedCursor {
		t.Errorf("%d != %d", win.cursor, expectedCursor)
	}
	expectedLine := "abcd"
	if win.lines[0] != expectedLine {
		t.Errorf("%d != %s", win.cursor, expectedLine)
	}
}

func TestRemovalBetweenLines(t *testing.T) {
	win := createWindowFromString("abcd\r\nqwer", 80, 200)
	cursorDown(win)
	enterInsertMode(win)
	removeUnderCursor(win)
	enterNormalMode(win)
	expectedCursor := CursorPos{0, 3}
	if *win.cursor != expectedCursor {
		t.Errorf("%d != %d", win.cursor, expectedCursor)
	}
	expectedLine := "abcdqwer"
	if win.lines[0] != expectedLine {
		t.Errorf("%d != %s", win.cursor, expectedLine)
	}
}
