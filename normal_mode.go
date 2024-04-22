package main

import "github.com/gdamore/tcell/v2"

func enterNormalMode(win *Window) {
	content := win.content
	win.cursor = min(win.cursor, len(win.content)-1)
	cursor := win.cursor
	isOnNewLine := isNewLine(content[cursor])
	isAfterNewLine := cursor != 0 && isNewLine(content[cursor-1])
	if isOnNewLine && isAfterNewLine {
		win.cursor--
	}
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
