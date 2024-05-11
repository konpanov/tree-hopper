package main

import "github.com/gdamore/tcell/v2"

func enterNormalMode(buf *Buffer) {
	buf.cursor.char = max(min(buf.cursor.char-1, mostRight(buf)), 0)
	buf.mode = NormalMode
	buf.cursor_style = tcell.CursorStyleSteadyBlock
}

func handleNormalModeEvents(buf *Buffer, ev *tcell.EventKey) {
	if buf.mode != NormalMode {
		return
	}
	switch ev.Key() {
	case tcell.KeyCtrlS:
		writeFile(buf)
	case tcell.KeyRune:
		switch ev.Rune() {
		case 'i':
			enterInsertMode(buf)
		case 'a':
			enterInsertMode(buf)
			cursorRight(buf)
		case 'v':
			enterVisualMode(buf)
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
