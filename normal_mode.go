package main

import "github.com/gdamore/tcell/v2"

func enterNormalMode(win *Window) {
	win.cursor.char = max(min(win.cursor.char-1, mostRight(win)), 0)
	win.mode = NormalMode
	win.cursor_style = tcell.CursorStyleSteadyBlock
}

func handleNormalModeEvents(win *Window, ev *tcell.EventKey) {
	if win.mode != NormalMode {
		return
	}
	switch ev.Key() {
	case tcell.KeyCtrlS:
		writeFile(win)
	case tcell.KeyRune:
		switch ev.Rune() {
		case 'i':
			enterInsertMode(win)
		case 'a':
			enterInsertMode(win)
			cursorRight(win)
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
}
