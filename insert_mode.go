package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

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
	case tcell.KeyEnter:
		splitLineUnderCursor(win)
	case tcell.KeyRune:
		insertUnderCursor(win, string(ev.Rune()))

	}
}

func removeUnderCursor(win *Window) {
	char := win.cursor.char
	line := win.cursor.line
	if char == 0 && line == 0 {
		return
	}
	if char == 0 {
		win.cursor.line = win.cursor.line - 1
		win.cursor.char = mostRight(win)
		mergeLines(win, win.cursor.line)
	} else {
		var sb strings.Builder
		line := win.lines[win.cursor.line]
		char := win.cursor.char
		sb.WriteString(line[:char-1])
		sb.WriteString(line[char:])
		win.lines[win.cursor.line] = sb.String()
		cursorLeft(win)
	}
}

func insertUnderCursor(win *Window, content string) {
	char := win.cursor.char
	line := win.lines[win.cursor.line]
	var sb strings.Builder
	sb.WriteString(line[:char])
	sb.WriteString(content)
	sb.WriteString(line[char:])
	win.lines[win.cursor.line] = sb.String()
	cursorRight(win)
}

func mergeLines(win *Window, line int) {
	var sb strings.Builder
	sb.WriteString(win.lines[line])
	sb.WriteString(win.lines[line+1])
	win.lines[line] = sb.String()
	win.lines = append(win.lines[:line+1], win.lines[line+2:]...)

}

func splitLineUnderCursor(win *Window) {
	line := win.cursor.line
	char := win.cursor.char
	if line < len(win.lines)-1 {
		win.lines = append(win.lines[:line+1], win.lines[line:]...)
	} else {
		win.lines = append(win.lines, "")
	}
	win.lines[line+1] = win.lines[line][char:]
	win.lines[line] = win.lines[line][:char]
	win.cursor.char = 0
	win.cursor.line = line + 1
}
