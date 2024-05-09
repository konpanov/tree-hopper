package main

import (
	"testing"
)

func TestCreateWindowFromStringSize(t *testing.T) {
	win := createWindowFromString("qwe\r\nasd", 80, 200)
	assertIntEqual(t, len(win.lines), 2)
}

func TestCreateWindowFromFileSize(t *testing.T) {
	win := createWindow("assets/tests/42lines.txt", 80, 40)
	assertStringEqual(t, win.lines[0], "1")
	assertStringEqual(t, win.lines[len(win.lines)-1], "42")
	assertIntEqual(t, len(win.lines), 42)
}

func TestReadFileSize(t *testing.T) {
	content := readFile("assets/tests/42lines.txt")
	assertByteEqual(t, content[len(content)-1], '2')
}
