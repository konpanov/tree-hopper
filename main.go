package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

type Mode int8

const (
	NormalMode Mode = iota
	InsertMode Mode = iota
)

type Window struct {
	filename     string
	content      string
	cursor_style tcell.CursorStyle
	cursor       int
	quiting      bool
	mode         Mode
}

func main() {
	filename := "main.go"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}
	win := createWindow(filename)
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	defer quit(screen, win)
	for !win.quiting {
		drawWindow(screen, win)
		handleEvents(screen.PollEvent(), win)
	}
}

func createWindow(filename string) *Window {
	window := &Window{
		filename:     filename,
		content:      readFile(filename),
		cursor:       0,
		cursor_style: tcell.CursorStyleSteadyBlock,
		quiting:      false,
		mode:         NormalMode,
	}
	return window
}

func quit(screen tcell.Screen, win *Window) {
	maybePanic := recover()
	screen.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}
