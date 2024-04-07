package main

type Cursor struct {
	col, row  int
	pos       int
	linestart int
	linewidth int
}

func cursorLeft(win *Window) {
	if !isLineStart(win.content, win.cursor.pos) {
		win.cursor.pos--
	}
}
func cursorRight(win *Window) {
	if !isLineEnd(win.content, win.cursor.pos) {
		win.cursor.pos++
	}
}

func cursorUp(win *Window) {
	for !isLineStart(win.content, win.cursor.pos) {
		win.cursor.pos--
	}
	if win.cursor.pos != 0 {
		win.cursor.pos--
		if win.content[win.cursor.pos] == '\n' {
			win.cursor.pos--
		}
		if win.content[win.cursor.pos] == '\r' {
			win.cursor.pos--
		}
	}
}

func cursorDown(win *Window) {
	for !isLineEnd(win.content, win.cursor.pos) {
		win.cursor.pos++
	}
	if !isEOF(win.content, win.cursor.pos) {
		win.cursor.pos++
		if win.content[win.cursor.pos] == '\r' {
			win.cursor.pos++
		}
		if win.content[win.cursor.pos] == '\n' {
			win.cursor.pos++
		}
	}
}

func isEOF(content string, pos int) bool {
	return pos == len(content)-1
}

func isLineStart(content string, pos int) bool {
	return pos == 0 || isNewLine(content[pos-1])
}
func isLineEnd(content string, pos int) bool {
	return pos == len(content)-1 || isNewLine(content[pos+1])
}

func isNewLine(r byte) bool {
	return r == '\r' || r == '\n'
}
