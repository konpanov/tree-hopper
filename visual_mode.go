package main

import (
	"github.com/gdamore/tcell/v2"
)

func enterVisualMode(buf *Buffer) {
	buf.mode = VisualMode
	buf.cursor_style = tcell.CursorStyleSteadyBlock
	buf.visualOrigin.line = buf.cursor.line
	buf.visualOrigin.char = buf.cursor.char
}

func handleVisualModeEvents(buf *Buffer, ev *tcell.EventKey) {
	if buf.mode != VisualMode {
		return
	}
	switch ev.Key() {
	case tcell.KeyEsc:
		enterNormalMode(buf)
	case tcell.KeyRune:
		switch ev.Rune() {
		case 'h':
			cursorLeft(buf)
		case 'j':
			cursorDown(buf)
		case 'k':
			cursorUp(buf)
		case 'l':
			cursorRight(buf)
		}
	}
}
