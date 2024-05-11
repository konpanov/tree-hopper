package main

import (
	"github.com/gdamore/tcell/v2"
)

func enterVisualMode(win *Window) {
	win.mode = VisualMode
	win.cursor_style = tcell.CursorStyleSteadyBlock
	win.visualOrigin.line = win.cursor.line
	win.visualOrigin.char = win.cursor.char
}

func handleVisualModeEvents(win *Window, ev *tcell.EventKey) {
	if win.mode != VisualMode {
		return
	}
	switch ev.Key() {
	case tcell.KeyEsc:
		enterNormalMode(win)
	case tcell.KeyRune:
		switch ev.Rune() {
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
