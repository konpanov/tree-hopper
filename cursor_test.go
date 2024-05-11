package main

import (
	"testing"
)

func TestCursorRight(t *testing.T) {
	buf := createBufferFromString("qwe\r\nasd", 80, 200)
	cursorRight(buf)
	expected := CursorPos{0, 1}
	if *buf.cursor != expected {
		t.Errorf("%d != %d", buf.cursor, expected)
	}
}

func TestCursorRightDoesNotSkipLines(t *testing.T) {
	buf := createBufferFromString("qwe\r\nasd", 80, 200)
	cursorRight(buf)
	cursorRight(buf)
	cursorRight(buf)
	expected := CursorPos{0, 2}
	if *buf.cursor != expected {
		t.Errorf("%d != %d", buf.cursor, expected)
	}
}

func TestCursorLeftDoesNotSkipLines(t *testing.T) {
	buf := createBufferFromString("qwe\r\nasd", 80, 200)
	cursorDown(buf)
	cursorRight(buf)
	cursorLeft(buf)
	cursorLeft(buf)
	cursorLeft(buf)
	cursorLeft(buf)
	expected := CursorPos{1, 0}
	if *buf.cursor != expected {
		t.Errorf("%d != %d", buf.cursor, expected)
	}
}

func TestCursorDown(t *testing.T) {
	buf := createBufferFromString("qwe\r\nasd", 80, 200)
	cursorDown(buf)
	expected := CursorPos{1, 0}
	if *buf.cursor != expected {
		t.Errorf("%d != %d", buf.cursor, expected)
	}
}

func TestCursorDownAndUp(t *testing.T) {
	buf := createBufferFromString("qwe\r\nasd", 80, 200)
	cursorDown(buf)
	cursorUp(buf)
	expected := CursorPos{0, 0}
	if *buf.cursor != expected {
		t.Errorf("%d != %d", buf.cursor, expected)
	}
}

func TestCursorDownUpDown(t *testing.T) {
	buf := createBufferFromString("qwe\r\nasd", 80, 200)
	cursorDown(buf)
	cursorUp(buf)
	cursorDown(buf)
	expected := CursorPos{1, 0}
	if *buf.cursor != expected {
		t.Errorf("%d != %d", buf.cursor, expected)
	}
}

func TestSlideWindowOnDown(t *testing.T) {
	buf := createBufferFromString(createNLines(100), 80, 40)
	assertIntEqual(t, buf.topLine, 0)
	for i := 0; i < 40; i++ {
		cursorDown(buf)
	}
	assertIntEqual(t, buf.topLine, 1)
}

func TestSlideWindowUp(t *testing.T) {
	buf := createBufferFromString(createNLines(100), 80, 40)
	for i := 0; i < 60; i++ {
		cursorDown(buf)
	}
	assertIntEqual(t, buf.topLine, 21)
	for i := 0; i < 60; i++ {
		cursorDown(buf)
	}
}

func TestCursorStopAtLastLine(t *testing.T) {
	buf := createBufferFromFile("assets/tests/42lines.txt", 80, 40)
	for i := 0; i < 150; i++ {
		cursorDown(buf)
	}
	assertIntEqual(t, buf.cursor.line, 41)
}
