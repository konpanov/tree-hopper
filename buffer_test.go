package main

import (
	"testing"
)

func TestCreateBufferFromStringSize(t *testing.T) {
	win := createBufferFromString("qwe\r\nasd", 80, 200)
	assertIntEqual(t, len(win.lines), 2)
}

func TestCreateBufferFromFileSize(t *testing.T) {
	win := createBufferFromFile("assets/tests/42lines.txt", 80, 40)
	assertStringEqual(t, win.lines[0], "1")
	assertStringEqual(t, win.lines[len(win.lines)-1], "42")
	assertIntEqual(t, len(win.lines), 42)
}

func TestReadFileSize(t *testing.T) {
	content := readFile("assets/tests/42lines.txt")
	assertByteEqual(t, content[len(content)-1], '2')
}
