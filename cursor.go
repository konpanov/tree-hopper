package main

func cursorLeft(win *Window) {
	if !isLineStart(win.content, win.cursor) {
		win.cursor--
	}
}
func cursorRight(win *Window) {
	if !isLineEnd(win.content, win.cursor) {
		win.cursor++
	}
}

func cursorUp(win *Window) {
	for !isLineStart(win.content, win.cursor) {
		win.cursor--
	}
	if win.cursor!= 0 {
		win.cursor--
		if win.content[win.cursor] == '\n' {
			win.cursor--
		}
		if win.content[win.cursor] == '\r' {
			win.cursor--
		}
	}
}

func cursorDown(win *Window) {
	for !isLineEnd(win.content, win.cursor) {
		win.cursor++
	}
	if !isEOF(win.content, win.cursor) {
		win.cursor++
		if win.content[win.cursor] == '\r' {
			win.cursor++
		}
		if win.content[win.cursor] == '\n' {
			win.cursor++
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
