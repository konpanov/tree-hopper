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
		t.Errorf("%s != %s", win.lines[0], expectedLine)
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
		t.Errorf("%s != %s", win.lines[0], expectedLine)
	}
}

func TestRemovalBetweenLines(t *testing.T) {
	win := createWindowFromString("abcd\r\nqwer", 80, 200)
	win.cursor.line = 1
	removeUnderCursor(win)
	expectedCursor := CursorPos{0, 3}
	if *win.cursor != expectedCursor {
		t.Errorf("%d != %d", win.cursor, expectedCursor)
	}
	expectedLine := "abcdqwer"
	if win.lines[0] != expectedLine {
		t.Errorf("%s != %s", win.lines[0], expectedLine)
	}
}

func TestSplitLine(t *testing.T) {
	win := createWindowFromString("abcdqwer", 80, 200)
	win.cursor.char = 4
	splitLineUnderCursor(win)
	expectedCursor := CursorPos{1, 0}
	if *win.cursor != expectedCursor {
		t.Errorf("%d != %d", win.cursor, expectedCursor)
	}
	if win.lines[0] != "abcd" {
		t.Errorf("%s != %s", win.lines[0], "abcd")
	}
	if win.lines[1] != "qwer" {
		t.Errorf("%s != %s", win.lines[1], "qwer")
	}
}

func TestSplitEmptyLine(t *testing.T) {
	win := createWindowFromString("", 80, 200)
	splitLineUnderCursor(win)
	expectedCursor := CursorPos{1, 0}
	if *win.cursor != expectedCursor {
		t.Errorf("%d != %d", win.cursor, expectedCursor)
	}
	if win.lines[0] != "" {
		t.Errorf("%s != %s", win.lines[0], "")
	}
	if win.lines[1] != "" {
		t.Errorf("%s != %s", win.lines[1], "")
	}
}

func TestSplitNotLastLine(t *testing.T) {
	win := createWindowFromString("qwe\r\nasd\r\nzxc", 80, 200)
	win.cursor = &CursorPos{1, 1}
	splitLineUnderCursor(win)
	expectedCursor := CursorPos{2, 0}
	if *win.cursor != expectedCursor {
		t.Errorf("%d != %d", win.cursor, expectedCursor)
	}
	if win.lines[0] != "qwe" {
		t.Errorf("%s != %s", win.lines[0], "qwe")
	}
	if win.lines[1] != "a" {
		t.Errorf("%s != %s", win.lines[0], "a")
	}
	if win.lines[2] != "sd" {
		t.Errorf("%s != %s", win.lines[1], "sd")
	}
	if win.lines[3] != "zxc" {
		t.Errorf("%s != %s", win.lines[0], "zxc")
	}
}
