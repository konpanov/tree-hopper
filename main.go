package main

import (
	"log"
	"os"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type Mode int8

const (
	NormalMode Mode = iota
	InsertMode Mode = iota
)

type CursorPos struct {
	line int
	char int
}

type Window struct {
	filename     string
	content      string
	lines        []string
	cursor_style tcell.CursorStyle
	cursor       *CursorPos
	quiting      bool
	mode         Mode
	hieght       int
	width        int
	newLineChar  string
}

func main() {
	filename := "main.go"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	w, h := screen.Size()
	win := createWindow(filename, w, h)
	defer quit(screen, win)
	for !win.quiting {
		drawWindow(screen, win)
		handleEvents(screen.PollEvent(), win)
	}
}

func createWindowFromString(content string, width, height int) *Window {
	newLineChar := "\r\n"
	window := &Window{
		filename:     "",
		content:      content,
		lines:        strings.Split(content, newLineChar),
		cursor:       &CursorPos{0, 0},
		cursor_style: tcell.CursorStyleSteadyBlock,
		quiting:      false,
		mode:         NormalMode,
		width:        width,
		hieght:       height,
		newLineChar:  newLineChar,
	}
	return window
}

func createWindow(filename string, width, height int) *Window {
	content := readFile(filename)
	window := createWindowFromString(content, width, height)
	window.filename = filename
	return window
}

func windowContent(win *Window) string {
	return strings.Join(win.lines, win.newLineChar)
}

func quit(screen tcell.Screen, win *Window) {
	maybePanic := recover()
	screen.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}
