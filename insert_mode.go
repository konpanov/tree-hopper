package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
)

func enterInsertMode(buf *Buffer) {
	buf.mode = InsertMode
	buf.cursor_style = tcell.CursorStyleBlinkingBar
}

func handleInsertModeEvents(buf *Buffer, ev *tcell.EventKey) {
	if buf.mode != InsertMode {
		return
	}
	switch ev.Key() {
	case tcell.KeyEsc:
		enterNormalMode(buf)
	case tcell.KeyBS:
		removeUnderCursor(buf)
	case tcell.KeyEnter:
		splitLineUnderCursor(buf)
	case tcell.KeyRune:
		insertUnderCursor(buf, string(ev.Rune()))

	}
}

func removeUnderCursor(buf *Buffer) {
	char := buf.cursor.char
	line := buf.cursor.line
	if char == 0 && line == 0 {
		return
	}
	if char == 0 {
		buf.cursor.line = buf.cursor.line - 1
		buf.cursor.char = mostRight(buf)
		mergeLines(buf, buf.cursor.line)
	} else {
		var sb strings.Builder
		line := buf.lines[buf.cursor.line]
		char := buf.cursor.char
		sb.WriteString(line[:char-1])
		sb.WriteString(line[char:])
		buf.lines[buf.cursor.line] = sb.String()
		cursorLeft(buf)
	}
}

func insertUnderCursor(buf *Buffer, content string) {
	char := buf.cursor.char
	line := buf.lines[buf.cursor.line]
	var sb strings.Builder
	sb.WriteString(line[:char])
	sb.WriteString(content)
	sb.WriteString(line[char:])
	buf.lines[buf.cursor.line] = sb.String()
	cursorRight(buf)
}

func mergeLines(buf *Buffer, line int) {
	var sb strings.Builder
	sb.WriteString(buf.lines[line])
	sb.WriteString(buf.lines[line+1])
	buf.lines[line] = sb.String()
	buf.lines = append(buf.lines[:line+1], buf.lines[line+2:]...)

}

func splitLineUnderCursor(buf *Buffer) {
	line := buf.cursor.line
	char := buf.cursor.char
	if line < len(buf.lines)-1 {
		buf.lines = append(buf.lines[:line+1], buf.lines[line:]...)
	} else {
		buf.lines = append(buf.lines, "")
	}
	buf.lines[line+1] = buf.lines[line][char:]
	buf.lines[line] = buf.lines[line][:char]
	buf.cursor.char = 0
	buf.cursor.line = line + 1
}
