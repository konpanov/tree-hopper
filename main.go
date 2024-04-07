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
	filename string
	content  string
	screen   tcell.Screen
	cursor   int
	quiting  bool
	mode     Mode
}

func main() {
	win := createWindow("main.go")
	defer quit(win)
	for !win.quiting {
		drawWindow(win)
		handleEvents(win)
	}
}

func createWindow(filename string) *Window {
	s, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	window := &Window{
		filename: filename,
		content:  string(dat),
		screen:   s,
		cursor:   0,
		quiting:  false,
		mode:     NormalMode,
	}
	s.SetCursorStyle(tcell.CursorStyleSteadyBlock)
	s.Clear()
	return window
}

func quit(win *Window) {
	maybePanic := recover()
	win.screen.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}

func drawWindow(win *Window) {
	col, row := 0, 0
	for i, r := range win.content {
		if i == win.cursor {
			win.screen.ShowCursor(col, row)
		}
		switch r {
		case '\r':
		case '\n':
			col = 0
			row++
		case '\t':
			_col := col
			for ; col < _col+8; col++ {
				SetContent(win, col, row, r)
			}
		default:
			SetContent(win, col, row, r)
			col++
		}
	}
	win.screen.Show()
}

func SetContent(win *Window, col, row int, r rune) {
	win.screen.SetContent(col, row, r, nil, tcell.Style{})
}

func handleEvents(win *Window) {
	ev := win.screen.PollEvent()

	switch ev := ev.(type) {
	case *tcell.EventKey:
		handleQuitEvent(win, ev)
		handleInsertModeEvents(win, ev)
		handleNormalModeEvents(win, ev)
	}
}

func handleNormalModeEvents(win *Window, ev *tcell.EventKey) {
	if win.mode != NormalMode {
		return
	}
	switch ev.Rune() {
	case 'i':
		enterInsertMode(win)
	case 'h':
		cursorLeft(win)
	case 'j':
		cursorDown(win)
	case 'k':
		cursorUp(win)
	case 'l':
		cursorRight(win)
	}
}

func handleInsertModeEvents(win *Window, ev *tcell.EventKey) {
	if win.mode != InsertMode {
		return
	}
	switch ev.Key() {
	case tcell.KeyEsc:
		enterNormalMode(win)
	case tcell.KeyRune:

	}
}

func handleQuitEvent(window *Window, ev *tcell.EventKey) {
	window.quiting = ev.Key() == tcell.KeyCtrlC
}

func enterInsertMode(win *Window) {
	win.mode = InsertMode
	win.screen.SetCursorStyle(tcell.CursorStyleBlinkingBar)
}

func enterNormalMode(win *Window) {
	win.mode = NormalMode
	win.screen.SetCursorStyle(tcell.CursorStyleSteadyBlock)
}
