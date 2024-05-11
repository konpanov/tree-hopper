package main

import (
	"testing"
)

func TestInsertSymbol(t *testing.T) {
	buf := createBufferFromString("", 80, 200)
	enterInsertMode(buf)
	content := "h"
	insertUnderCursor(buf, content)
	expected := content
	if buf.lines[0] != expected {
		t.Errorf("%d != %s", buf.cursor, expected)
	}
}

func TestCursorPosAfterInsert(t *testing.T) {
	buf := createBufferFromString("", 80, 200)
	enterInsertMode(buf)
	content := "h"
	insertUnderCursor(buf, content)
	expected := CursorPos{0, 1}
	if *buf.cursor != expected {
		t.Errorf("%d != %d", buf.cursor, expected)
	}
}

func TestCursorPosAfterMultipleInserts(t *testing.T) {
	buf := createBufferFromString("", 80, 200)
	enterInsertMode(buf)
	content := "h"
	insertUnderCursor(buf, content)
	insertUnderCursor(buf, content)
	insertUnderCursor(buf, content)
	expected := CursorPos{0, 3}
	if *buf.cursor != expected {
		t.Errorf("%d != %d", buf.cursor, expected)
	}
}

func TestCursorPosInNormalModeAfterInsert(t *testing.T) {
	buf := createBufferFromString("", 80, 200)
	enterInsertMode(buf)
	content := "h"
	insertUnderCursor(buf, content)
	enterNormalMode(buf)
	expected := CursorPos{0, 0}
	if *buf.cursor != expected {
		t.Errorf("%d != %d", buf.cursor, expected)
	}
}

func TestRemoval(t *testing.T) {
	buf := createBufferFromString("abcd", 80, 200)
	cursorRight(buf)
	cursorRight(buf)
	cursorRight(buf)
	enterInsertMode(buf)
	removeUnderCursor(buf)
	enterNormalMode(buf)
	expectedCursor := CursorPos{0, 1}
	if *buf.cursor != expectedCursor {
		t.Errorf("%d != %d", buf.cursor, expectedCursor)
	}
	expectedLine := "abd"
	if buf.lines[0] != expectedLine {
		t.Errorf("%s != %s", buf.lines[0], expectedLine)
	}
}

func TestRemovalOnFirstSymbol(t *testing.T) {
	buf := createBufferFromString("abcd", 80, 200)
	enterInsertMode(buf)
	removeUnderCursor(buf)
	enterNormalMode(buf)
	expectedCursor := CursorPos{0, 0}
	if *buf.cursor != expectedCursor {
		t.Errorf("%d != %d", buf.cursor, expectedCursor)
	}
	expectedLine := "abcd"
	if buf.lines[0] != expectedLine {
		t.Errorf("%s != %s", buf.lines[0], expectedLine)
	}
}

func TestRemovalBetweenLines(t *testing.T) {
	buf := createBufferFromString("abcd\r\nqwer", 80, 200)
	buf.cursor.line = 1
	removeUnderCursor(buf)
	expectedCursor := CursorPos{0, 3}
	if *buf.cursor != expectedCursor {
		t.Errorf("%d != %d", buf.cursor, expectedCursor)
	}
	expectedLine := "abcdqwer"
	if buf.lines[0] != expectedLine {
		t.Errorf("%s != %s", buf.lines[0], expectedLine)
	}
}

func TestSplitLine(t *testing.T) {
	buf := createBufferFromString("abcdqwer", 80, 200)
	buf.cursor.char = 4
	splitLineUnderCursor(buf)
	expectedCursor := CursorPos{1, 0}
	if *buf.cursor != expectedCursor {
		t.Errorf("%d != %d", buf.cursor, expectedCursor)
	}
	if buf.lines[0] != "abcd" {
		t.Errorf("%s != %s", buf.lines[0], "abcd")
	}
	if buf.lines[1] != "qwer" {
		t.Errorf("%s != %s", buf.lines[1], "qwer")
	}
}

func TestSplitEmptyLine(t *testing.T) {
	buf := createBufferFromString("", 80, 200)
	splitLineUnderCursor(buf)
	expectedCursor := CursorPos{1, 0}
	if *buf.cursor != expectedCursor {
		t.Errorf("%d != %d", buf.cursor, expectedCursor)
	}
	if buf.lines[0] != "" {
		t.Errorf("%s != %s", buf.lines[0], "")
	}
	if buf.lines[1] != "" {
		t.Errorf("%s != %s", buf.lines[1], "")
	}
}

func TestSplitNotLastLine(t *testing.T) {
	buf := createBufferFromString("qwe\r\nasd\r\nzxc", 80, 200)
	buf.cursor = &CursorPos{1, 1}
	splitLineUnderCursor(buf)
	expectedCursor := CursorPos{2, 0}
	if *buf.cursor != expectedCursor {
		t.Errorf("%d != %d", buf.cursor, expectedCursor)
	}
	if buf.lines[0] != "qwe" {
		t.Errorf("%s != %s", buf.lines[0], "qwe")
	}
	if buf.lines[1] != "a" {
		t.Errorf("%s != %s", buf.lines[0], "a")
	}
	if buf.lines[2] != "sd" {
		t.Errorf("%s != %s", buf.lines[1], "sd")
	}
	if buf.lines[3] != "zxc" {
		t.Errorf("%s != %s", buf.lines[0], "zxc")
	}
}
