package main

import (
	"testing"
)

func TestVisualMovment(t *testing.T) {
	buf := createBufferFromString("qwe\r\nasd", 80, 200)
	cursorRight(buf)
	enterVisualMode(buf)
	cursorDown(buf)

	vo := buf.visualOrigin

	if vo.char != 1 {
		temp := "Incorrect visual origin char %d != %d"
		t.Errorf(temp, vo.char, 1)
	}
	if vo.line != 0 {
		temp := "Incorrect visual origin line %d != %d"
		t.Errorf(temp, vo.char, 0)
	}
	if buf.cursor.char != 1 {
		temp := "Incorrect cursor char %d != %d"
		t.Errorf(temp, buf.cursor.char, 0)
	}
	if buf.cursor.line != 1 {
		temp := "Incorrect cursor line %d != %d"
		t.Errorf(temp, buf.cursor.line, 1)
	}
}
