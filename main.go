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

type Cursor struct {
	col, row int
}

type Window struct {
	filename string
	content  string
	screen   tcell.Screen
	cursor   *Cursor
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
		cursor:   &Cursor{0, 0},
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
	for _, rune := range win.content {
		if rune == '\n' {
			col = 0
			row++
			continue
		}
		if rune == 9 {
			for i := 0; i < 7; i++ {
				win.screen.SetContent(col, row, rune, nil, tcell.Style{})
				col++
			}
		}
		win.screen.SetContent(col, row, rune, nil, tcell.Style{})
		col++
	}
	win.screen.ShowCursor(win.cursor.col, win.cursor.row)
	win.screen.Show()
}

func handleEvents(win *Window) {
	ev := win.screen.PollEvent()

	switch ev := ev.(type) {
	case *tcell.EventKey:
		handleEscapeEvent(win, ev)
		handleModeEvents(win, ev)
	}
}

func handleModeEvents(win *Window, ev *tcell.EventKey) {
	handleInsertModeEvents(win, ev)
	handleNormalModeEvents(win, ev)
}

func handleNormalModeEvents(win *Window, ev *tcell.EventKey) {
	if win.mode != NormalMode {
		return
	}
	switch ev.Rune() {
	case 'i':
		enterInsertMode(win)
	case 'h':
		win.cursor.col = max(0, win.cursor.col-1)
	case 'j':
		_, height := win.screen.Size()
		win.cursor.row = min(height-1, win.cursor.row+1)
	case 'k':
		win.cursor.row = max(0, win.cursor.row-1)
	case 'l':
		width, _ := win.screen.Size()
		win.cursor.col = min(width-1, win.cursor.col+1)
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

func handleEscapeEvent(window *Window, ev *tcell.EventKey) {
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
