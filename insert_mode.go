package main

import "github.com/gdamore/tcell/v2"

func enterInsertMode(win *Window) {
	win.mode = InsertMode
	win.cursor_style = tcell.CursorStyleBlinkingBar
}

func handleInsertModeEvents(win *Window, ev *tcell.EventKey) {
	if win.mode != InsertMode {
		return
	}
	switch ev.Key() {
	case tcell.KeyEsc:
		enterNormalMode(win)
	case tcell.KeyBS:
		removeUnderCursor(win)
	case tcell.KeyRune:
		insertUndercursor(win, string(ev.Rune()))

	}
}

func removeUnderCursor(win *Window) {
	if win.cursor > 1 && win.content[win.cursor-2] == '\r' && win.content[win.cursor-1] == '\n' {
		win.content = win.content[:win.cursor-2] + win.content[win.cursor:]
		cursorLeft(win)
		cursorLeft(win)
	} else {
		win.content = win.content[:win.cursor-1] + win.content[win.cursor:]
		cursorLeft(win)
	}
}

func insertUndercursor(win *Window, content string) {
	win.content = win.content[:win.cursor] + string(content) + win.content[win.cursor:]
	win.cursor++
}
